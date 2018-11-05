package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
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
	if c.GetAuthor().Username == "SIREL" {
		c.WE(fmt.Errorf("User SIREL can't reserve"), 403)
	}

	lc := ReservationToCreate{}
	c.ReadObjectInBody("reservation", &lc, true)

	li := models.ReservationInfo{}
	li.ID = lc.LocalID
	li.UserID = c.GetAuthor().ID
	li.LocalID = lc.LocalID
	li.ActivityName = lc.ActivityName
	li.ActivityDescription = lc.ActivityDescription
	li.BeginTime = lc.BeginTime
	li.EndTime = lc.EndTime
	c.Validate(li)

	l, me, e := app.Model().AddReservation(li)
	if e != nil && me {
		c.WE(e, 400)
	}
	c.WE(e, 500)
	return l
}

// AcceptReservation ...
func (c *BaseReservationsController) AcceptReservation() {
	r := c.loadReservation()
	c.isLocalAdmin(r.LocalID, c.GetAuthor().ID)

	// Notificate to user
	local, e := r.Local()
	c.WE(e, 500)

	year, month, day := r.BeginTime.Date()
	bh, bm, _ := r.BeginTime.Clock()
	eh, em, _ := r.EndTime.Clock()

	e = app.Model().NotificateToUser(r.UserID,
		fmt.Sprintf("Su reservacion en el local %s con fecha %d/%d/%d %d:%d - %d:%d fue aceptada",
			local.Name, year, month, day, bh, bm, eh, em))
	c.WE(e, 500)

	r.Pending = false
	e = r.Update()
	c.WE(e, 500)
}

// RefuseReservation ...
func (c *BaseReservationsController) RefuseReservation() {
	r := c.loadReservation()
	c.isLocalAdmin(r.LocalID, c.GetAuthor().ID)

	// Notificate to user
	local, e := r.Local()
	c.WE(e, 500)

	year, month, day := r.BeginTime.Date()
	bh, bm, _ := r.BeginTime.Clock()
	eh, em, _ := r.EndTime.Clock()

	e = app.Model().NotificateToUser(r.UserID,
		fmt.Sprintf("Su reservacion en el local %s con fecha %d/%d/%d %d:%d - %d:%d fue denegada",
			local.Name, year, month, day, bh, bm, eh, em))

	c.WE(e, 500)

	e = app.Model().Delete(r)
	c.WE(e, 500)
}

// List ...
func (c *BaseReservationsController) List() *models.ReservationCollection {
	limit, offset, orderby, desc := c.ReadPagOrder()
	userID := c.ReadInt("user_id")
	localID := c.ReadInt("local_id")
	confirmed := c.ReadBool("confirmed")
	pending := c.ReadBool("pending")
	sdate := c.ReadString("date")
	beego.Debug("before localAdminID")
	localAdminID := c.ReadInt("localAdminID")
	beego.Debug("after localAdminID")
	search := c.ReadString("search")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	rs, e := app.Model().GetReservations(search, userID, localID, confirmed,
		pending, date, localAdminID, limit, offset, orderby, desc)

	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return rs
}

// ReservationToCreate ...
type ReservationToCreate struct {
	LocalID             int
	ActivityName        string
	ActivityDescription string
	BeginTime           time.Time
	EndTime             time.Time
}

func (c *BaseReservationsController) loadReservation() *models.Reservation {
	r := app.Model().NewReservation()
	r.ID = *(c.ReadInt("reservation_id", true))
	e := r.Load()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return r
}

func (c *BaseReservationsController) isLocalAdmin(localID, userID int) {
	_, e := app.Model().GetLocalAdmin(localID, userID)
	if e == models.ErrNoRows {
		c.WE(fmt.Errorf("Forbbiden"), 403)
	}
	c.WE(e, 500)
}
