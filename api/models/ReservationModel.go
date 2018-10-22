package models

import (
	"time"

	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// ReservationInfo ...
type ReservationInfo struct {
	ID                  int
	UserID              int
	LocalID             int
	ActivityName        string
	ActivityDescription string
	BeginTime           time.Time
	EndTime             time.Time
	Confirmed           bool
	Pending             bool
}

// Reservation ...
type Reservation struct {
	ReservationInfo
	model Model

	local *Local
	user  *User
}

/////////////////////////////////////////////////////

// TableName ...
func (r *Reservation) TableName() string {
	return "reservation"
}

// AutoPKey ...
func (r *Reservation) AutoPKey() bool {
	return true
}

// PkeyName ...
func (r *Reservation) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (r *Reservation) PkeyValue() interface{} {
	return r.ID
}

// PkeyPointer ...
func (r *Reservation) PkeyPointer() interface{} {
	return &r.ID
}

// ColumnNames ...
func (r *Reservation) ColumnNames() []string {
	return []string{
		"user_id",
		"local_id",
		"activity_name",
		"activity_description",
		"begin_time",
		"end_time",
		"confirmed",
		"pending",
	}
}

// ColumnValues ...
func (r *Reservation) ColumnValues() []interface{} {
	return []interface{}{
		r.UserID,
		r.LocalID,
		r.ActivityName,
		r.ActivityDescription,
		r.BeginTime,
		r.EndTime,
		r.Confirmed,
		r.Pending,
	}
}

// ColumnPointers ...
func (r *Reservation) ColumnPointers() []interface{} {
	return []interface{}{
		&r.UserID,
		&r.LocalID,
		&r.ActivityName,
		&r.ActivityDescription,
		&r.BeginTime,
		&r.EndTime,
		&r.Confirmed,
		&r.Pending,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (r *Reservation) Update() error {
	return r.model.Update(r)
}

// Load ...
func (r *Reservation) Load() error {
	return r.model.Retrieve(r)
}

// Local ...
func (r *Reservation) Local() (*Local, error) {
	var e error
	if r.local == nil {
		r.local = r.model.NewLocal()
		r.local.ID = r.LocalID
		e = r.model.Retrieve(r.local)
	}
	return r.local, e
}

// User ...
func (r *Reservation) User() (*User, error) {
	var e error
	if r.user == nil {
		r.user = r.model.NewUser()
		r.user.ID = r.UserID
		e = r.model.Retrieve(r.user)
	}
	return r.user, e
}

///////////////////////////////////////////////////////////////////////////////////

// ReservationCollection ...
type ReservationCollection struct {
	model        Model
	Reservations *[]*Reservation
}

// NewObjectModel ...
func (c *ReservationCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewReservation()
}

// Add ...
func (c *ReservationCollection) Add() kmodel.ObjectModel {
	r := c.model.NewReservation()
	*(c.Reservations) = append(*(c.Reservations), r)
	return r
}

///////////////////////////////////////////////////////////////////////////////////

// ReservationModel ...
type ReservationModel interface {
	NewReservation() *Reservation
	NewReservationCollection() *ReservationCollection
	Reservations(limit, offset *int, orderby *string,
		orderDesc *bool) (*ReservationCollection, error)

	ReservationCustomModel
}

/////////////////////////////////////////////////////

// NewReservation ...
func (m *model) NewReservation() *Reservation {
	r := &Reservation{
		model: m,
	}
	return r
}

// NewReservationCollection ...
func (m *model) NewReservationCollection() *ReservationCollection {
	kk := make([]*Reservation, 0)
	return &ReservationCollection{
		model:        m,
		Reservations: &kk,
	}
}

func (m *model) Reservations(limit, offset *int, orderby *string,
	orderDesc *bool) (*ReservationCollection, error) {

	collection := m.NewReservationCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
