package controllers

import (
	"time"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

// BaseReservationsController ...
type BaseReservationsController struct {
	BaseController
}

// Show ...
func (c *BaseReservationsController) Show() *models.Reservation {
	r := app.Model().NewReservation()
	r.ID = *c.ReadInt("reservation_id", true)
	e := r.Load()

	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return r
}

// Create ...
func (c *BaseReservationsController) Create() *models.Reservation {
	l := app.Model().NewReservation()

	lc := ReservationToCreate{}
	c.ReadInputBody(&lc)

	l.ID = lc.ID
	l.UserID = c.GetAuthor().ID
	l.LocalID = lc.LocalID
	l.ActivityName = lc.ActivityName
	l.ActivityDescription = lc.ActivityDescription
	l.BeginTime = time.Now()
	l.EndTime = time.Now()
	l.Confirmed = true
	l.Pending = true

	c.Validate(l)

	e := app.Model().Create(l)
	c.WE(e, 500)
	return l
}

// List ...
func (c *BaseReservationsController) List() *models.ReservationCollection {
	limit := c.ReadInt("limit")
	offset := c.ReadInt("offset")
	orderby := c.ReadString("orderby")
	desc := c.ReadBool("orderDesc")
	userID := c.ReadInt("user_id")
	localID := c.ReadInt("local_id")
	confirmed := c.ReadBool("confirmed")
	pending := c.ReadBool("pending")
	sdate := c.ReadString("date", false)
	search := c.ReadString("search")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	rs, e := app.Model().GetReservations(search, userID, localID, confirmed,
		pending, date, limit, offset, orderby, desc)

	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return rs
}

// ReservationToCreate ...
type ReservationToCreate struct {
	ID                  int
	LocalID             int
	ActivityName        string
	ActivityDescription string
	BeginTime           time.Time
	EndTime             time.Time
}
