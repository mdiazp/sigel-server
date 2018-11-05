package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ReservationCustomModel ...
type ReservationCustomModel interface {
	GetReservations(search *string, userID, localID *int,
		confirmed *bool, pending *bool, date *Date, localAdminID *int,
		limit, offset *int, orderby *string, desc *bool) (*ReservationCollection, error)

	AddReservation(ri ReservationInfo) (*Reservation, bool, error)
	NewDate(s *string) (*Date, error)
}

func (m *model) GetReservations(search *string, userID, localID *int,
	confirmed *bool, pending *bool, date *Date, localAdminID *int,
	limit, offset *int, orderby *string, desc *bool) (*ReservationCollection, error) {

	where := ""
	if search != nil {
		if where != "" {
			where += " AND "
		}
		where += "reservation.activity_name like '%" + *search + "%'"
	}
	if userID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.user_id=%d", *userID)
	}
	if localID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.local_id=%d", *localID)
	}
	if confirmed != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.confirmed=%t", *confirmed)
	}
	if pending != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("reservation.pending=%t", *pending)
	}
	if date != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("extract(year from reservation.begin_time)=%d AND ", date.Year) +
			fmt.Sprintf("extract(month from reservation.begin_time)=%d AND ", date.Month) +
			fmt.Sprintf("extract(day from reservation.begin_time)=%d", date.Day)
	}

	if localAdminID != nil {
		if where != "" {
			where += " AND "
		}
		where += "reservation.local_id IN (" +
			fmt.Sprintf(
				"SELECT local_admin.local_id FROM local_admin "+
					"WHERE local_admin.user_id=%d", *localAdminID) +
			")"
	}

	hf := &where
	if where == "" {
		hf = nil
	}

	rs := m.NewReservationCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, rs)
	return rs, e
}

// AddReservation ...
func (m *model) AddReservation(ri ReservationInfo) (*Reservation, bool, error) {
	eLocalDontExist := fmt.Errorf("Don't axists any local with ID=%d", ri.LocalID)
	eInvalid := fmt.Errorf("Invalid reservation")
	eUnworked := fmt.Errorf("Non laboral date in this local")
	eConflictTime := fmt.Errorf("Conflict time with other reservations")

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

	// Validatinf if date is laboral
	if l.WorkingMonths[int(bt.Month())-1] == '0' ||
		l.WorkingWeekDays[int(bt.Weekday())-1] == '0' {
		return nil, true, eUnworked
	}
	if bt.Hour() < l.WorkingBeginTimeHours ||
		(bt.Hour() == l.WorkingBeginTimeHours && bt.Minute() < l.WorkingBeginTimeMinutes) ||
		et.Hour() > l.WorkingEndTimeHours ||
		(et.Hour() == l.WorkingEndTimeHours && et.Minute() > l.WorkingEndTimeMinutes) {
		return nil, true, eUnworked
	}

	// Validate that don't exist's conflict time with other reservations
	tmp := l.model.NewReservation()

	println("------------> bt =", ri.BeginTime.Format("2006-01-02 15:04:05"))
	println("------------> et =", ri.EndTime.Format("2006-01-02 15:04:05"))

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
	ri.Confirmed = l.WorkingWeekDays[int(bt.Weekday())-1] == '1'

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
	err := fmt.Errorf("date's format is invalid")
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
			return 0, fmt.Errorf("Invalid number")
		}
		x *= 10
		x += int(c) - int('0')
	}
	return x, nil
}
