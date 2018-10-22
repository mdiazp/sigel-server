package models

import (
	"fmt"
	"strconv"
	"strings"
)

// ReservationCustomModel ...
type ReservationCustomModel interface {
	GetReservations(search *string, userID, localID *int,
		confirmed *bool, pending *bool, date *Date,
		limit, offset *int, orderby *string, desc *bool) (*ReservationCollection, error)

	NewDate(s *string) (*Date, error)
}

func (m *model) GetReservations(search *string, userID, localID *int,
	confirmed *bool, pending *bool, date *Date,
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
	/*
		if date != nil {
			if where != "" {
				where += " AND "
			}
			where += fmt.Sprintf("Year(reservation.begin_time)=%d AND ", date.Year) +
				fmt.Sprintf("Month(reservation.begin_time)=%d AND ", date.Month) +
				fmt.Sprintf("Day(reservation.begin_time)=%d", date.Day)
		}
	*/

	hf := &where
	if where == "" {
		hf = nil
	}

	rs := m.NewReservationCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, rs)
	return rs, e
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
	if e != nil {
		d.Month, e = strconv.Atoi(x[1])
	}
	if e != nil {
		d.Day, e = strconv.Atoi(x[2])
	}
	if e == nil {
		err = nil
	}
	return &d, err
}
