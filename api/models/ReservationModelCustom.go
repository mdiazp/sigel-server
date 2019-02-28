package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
)

// ReservationCustomModel ...
type ReservationCustomModel interface {
	GetReservations(f ReservationFilter, limit, offset *int, orderby *string,
		desc *bool) (*ReservationCollection, error)
	GetReservationsCount(f ReservationFilter) (int, error)

	AddReservation(ri ReservationInfo) (*Reservation, bool, error)
	NewDate(s *string) (*Date, error)
}

// ReservationFilter ...
type ReservationFilter struct {
	UserID        *int
	LocalID       *int
	Confirmed     *bool
	Pending       *bool
	LocalAdminID  *int
	Search        *string
	Date          *Date
	NotBeforeDate *Date
}

func (m *model) GetReservations(f ReservationFilter, limit, offset *int, orderby *string,
	desc *bool) (*ReservationCollection, error) {

	hf := m.MakeReservationHorizontalFilter(f)

	if orderby == nil {
		tmp := "reservation.begin_time"
		orderby = &tmp
		tmp2 := true
		desc = &tmp2
	} else {
		*orderby = "reservation." + *orderby
	}

	rs := m.NewReservationCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, rs)
	return rs, e
}

func (m *model) GetReservationsCount(f ReservationFilter) (int, error) {
	hf := m.MakeReservationHorizontalFilter(f)

	o := m.NewReservation()
	count := 0
	e := m.RetrieveCount(hf, o, &count)
	return count, e
}

func (m *model) MakeReservationHorizontalFilter(f ReservationFilter) *string {
	where := ""
	if f.Search != nil {
		if where != "" {
			where += " AND "
		}
		where += "reservation.activity_name like '%" + *f.Search + "%'"
	}
	if f.UserID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.user_id=%d", *f.UserID)
	}
	if f.LocalID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.local_id=%d", *f.LocalID)
	}
	if f.Confirmed != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.confirmed=%t", *f.Confirmed)
	}
	if f.Pending != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.pending=%t", *f.Pending)
	}
	if f.Date != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("extract(year from reservation.begin_time)=%d AND ", f.Date.Year) +
			fmt.Sprintf("extract(month from reservation.begin_time)=%d AND ", f.Date.Month) +
			fmt.Sprintf("extract(day from reservation.begin_time)=%d", f.Date.Day)
	}
	if f.NotBeforeDate != nil {
		if where != "" {
			where += " AND "
		}
		where += "(" +
			fmt.Sprintf("extract(year from reservation.begin_time)>%d OR ", f.NotBeforeDate.Year) +
			fmt.Sprintf("(extract(year from reservation.begin_time)=%d AND extract(month from reservation.begin_time)>%d) OR", f.NotBeforeDate.Year, f.NotBeforeDate.Month) +
			fmt.Sprintf("(extract(year from reservation.begin_time)=%d AND extract(month from reservation.begin_time)=%d AND extract(day from reservation.begin_time)>=%d)", f.NotBeforeDate.Year, f.NotBeforeDate.Month, f.NotBeforeDate.Day) +
			")"
	}

	if f.LocalAdminID != nil {
		if where != "" {
			where += " AND "
		}
		where += "reservation.local_id IN (" +
			fmt.Sprintf(
				"SELECT local_admin.local_id FROM local_admin "+
					"WHERE local_admin.user_id=%d", *f.LocalAdminID) +
			")"
	}

	if where == "" {
		return nil
	}

	return &where
}

// AddReservation ...
func (m *model) AddReservation(ri ReservationInfo) (*Reservation, bool, error) {
	eLocalDontExist := fmt.Errorf("Local no encontrado")
	eInvalid := fmt.Errorf("Reservación incorrecta")
	eUnworked := fmt.Errorf("El local no está laborable es la fecha")
	eConflictTime := fmt.Errorf("Existe conflicto de tiempo con otras reservaciones")
	eMinDuration := fmt.Errorf("No se puede reservar por menos de 30 minutos")

	l := m.NewLocal()
	l.ID = ri.LocalID
	e := l.Load()

	if e != nil {
		return nil, true, eLocalDontExist
	}

	bt := ri.BeginTime
	et := ri.EndTime

	by, bm, bd := bt.Date()
	ey, em, ed := et.Date()

	if bt.After(et) || time.Now().After(bt) || by != ey || bm != em || bd != ed {
		return nil, true, eInvalid
	}

	// Validating if date is laboral
	if l.WorkingMonths[int(bt.Month())-1] == '0' ||
		l.WorkingWeekDays[int(bt.Weekday())] == '0' {
		return nil, true, eUnworked
	}

	if bt.Hour() < l.WorkingBeginTimeHours ||
		(bt.Hour() == l.WorkingBeginTimeHours && bt.Minute() < l.WorkingBeginTimeMinutes) ||
		et.Hour() > l.WorkingEndTimeHours ||
		(et.Hour() == l.WorkingEndTimeHours && et.Minute() > l.WorkingEndTimeMinutes) {
		return nil, true, eUnworked
	}

	// Validate that reservation interval has more than 30 minutes
	if (et.Hour()*60+et.Minute())-(bt.Hour()*60+bt.Minute())+1 < 30 {
		return nil, true, eMinDuration
	}

	// Validate that don't exists conflict time with other reservations
	tmp := l.model.NewReservation()

	e = l.model.RetrieveOne(tmp,
		"reservation.local_id=$1 AND NOT(reservation.end_time < $2 OR $3 < reservation.begin_time)",
		l.ID, ri.BeginTime.Format("2006-01-02 15:04:05"), ri.EndTime.Format("2006-01-02 15:04:05"))

	if e != ErrNoRows {
		if e != nil {
			return nil, false, e
		}
		return nil, true, eConflictTime
	}

	ri.LocalID = l.ID
	ri.Pending = true
	ri.Confirmed = l.WorkingWeekDays[int(bt.Weekday())] == '1'

	r := l.model.NewReservation()
	r.ReservationInfo = ri
	e = l.model.Create(r)

	return r, false, e
}

//Date ...
type Date struct {
	Year  int
	Month int
	Day   int
}

// NewDate return an object Date given a string with format yyyy-mm-dd
func (m *model) NewDate(s *string) (*Date, error) {
	err := fmt.Errorf("El formato de fecha no es válido")
	if s == nil {
		return nil, nil
	}

	x := strings.Split(*s, "-")
	if len(x) != 3 {
		return nil, err
	}

	var e error
	d := Date{}
	d.Year, e = strconv.Atoi(x[0])
	if e == nil {
		d.Month, e = strconv.Atoi(x[1])
	}
	if e == nil {
		d.Day, e = strconv.Atoi(x[2])
	}
	if e == nil {
		err = nil
	}
	return &d, err
}

func getnumber(s string) (int, error) {
	x := 0
	for _, c := range s {
		if c < '0' || '9' < c {
			return 0, fmt.Errorf("El valor debe contener solo dígitos")
		}
		x *= 10
		x += int(c) - int('0')
	}
	return x, nil
}

// Valid ...
func (r *ReservationInfo) Valid(v *validation.Validation) {
	validateNotEmptyString("activityName", r.ActivityName, v)
	validateNotEmptyString("activityDescription", r.ActivityDescription, v)
}
