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

	rc := ReservationToCreate{}
	c.ReadObjectInBody("reservation", &rc, true)

	ri := models.ReservationInfo{}
	ri.ID = rc.LocalID
	ri.UserID = c.GetAuthor().ID
	ri.LocalID = rc.LocalID
	ri.ActivityName = rc.ActivityName
	ri.ActivityDescription = rc.ActivityDescription
	ri.BeginTime = rc.BeginTime
	ri.EndTime = rc.EndTime
	c.Validate(ri)

	r, me, e := app.Model().AddReservation(ri)

	if e != nil && me {
		c.WE(e, 400)
	}
	c.WE(e, 500)
	return r
}

// Confirm ...
func (c *BaseController) Confirm() *models.Reservation {
	rID := c.ReadInt("reservationID", true)
	r := app.Model().NewReservation()
	r.ID = *rID
	e := r.Load()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	if c.GetAuthor().ID != r.UserID {
		c.WE(fmt.Errorf("Only the user that is author of a reservation can confirm it"), 403)
	}

	bt := r.BeginTime
	bt = bt.AddDate(0, 0, -1)

	st := time.Now()

	beego.Debug("bt.Year = ", bt.Year())
	beego.Debug("bt.Month = ", bt.Month())
	beego.Debug("bt.Day = ", bt.Day())
	beego.Debug("st.Year = ", st.Year())
	beego.Debug("st.Month = ", st.Month())
	beego.Debug("st.Day = ", st.Day())

	if bt.Year() != st.Year() || bt.Month() != st.Month() || bt.Day() != st.Day() {
		c.WE(fmt.Errorf("La reservación solo puede ser confirmada un día antes"), 400)
	}

	r.Confirmed = true
	e = r.Update()
	c.WE(e, 500)

	return r
}

// AcceptReservation ...
func (c *BaseReservationsController) AcceptReservation() {
	r := c.loadReservation()
	c.isLocalAdmin(r.LocalID, c.GetAuthor().ID)

	// Notificate to user
	local, e := r.Local()
	c.WE(e, 500)

	r.Pending = false
	e = r.Update()
	c.WE(e, 500)

	year, month, day := r.BeginTime.Date()
	bh, bm, _ := r.BeginTime.Clock()
	eh, em, _ := r.EndTime.Clock()

	// Notificate to user by email
	msg := fmt.Sprintf(
		"Su reservacion en el local %s con "+
			"fecha %d/%.2d/%.2d entre %.2d:%.2d y %.2d:%.2d fue aceptada",
		local.Name, year, month, day, bh, bm, eh, em,
	)
	notificateByEmail(r, msg)
}

// RefuseReservation ...
func (c *BaseReservationsController) RefuseReservation() {
	r := c.loadReservation()
	c.isLocalAdmin(r.LocalID, c.GetAuthor().ID)

	// Notificate to user
	local, e := r.Local()
	c.WE(e, 500)

	e = app.Model().Delete(r)
	c.WE(e, 500)

	year, month, day := r.BeginTime.Date()
	bh, bm, _ := r.BeginTime.Clock()
	eh, em, _ := r.EndTime.Clock()

	// Notificate to user by email
	msg := fmt.Sprintf(
		"Su reservacion en el local %s con "+
			"fecha %d/%.2d/%.2d entre %.2d:%.2d y %.2d:%.2d fue denegada",
		local.Name, year, month, day, bh, bm, eh, em,
	)
	notificateByEmail(r, msg)
}

func notificateByEmail(r *models.Reservation, msg string) {
	e := app.Model().NotificateToUser(r.UserID, msg)
	if e != nil {
		beego.Debug("Notification couldn't be saved for user: ", e.Error())
	}

	// Send Email
	u, e := app.Model().GetUserByID(r.UserID)
	if e == nil && u.SendNotificationsToEmail {
		e = app.GetMailSender().SendMail(u.Email, msg)
		if e != nil {
			beego.Debug("Notification couldn't be sended to user email: ", e.Error())
		}
	}
}

// List ...
func (c *BaseReservationsController) List() *models.ReservationCollection {
	limit, offset, orderby, desc := c.ReadPagOrder()
	userID := c.ReadInt("user_id")
	localID := c.ReadInt("local_id")
	confirmed := c.ReadBool("confirmed")
	pending := c.ReadBool("pending")
	localAdminID := c.ReadInt("localAdminID")
	search := c.ReadString("search")
	sdate := c.ReadString("date")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}
	sdate = c.ReadString("not_before_date")
	notBeforeDate, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	beego.Debug("desc !== nil  ======== ", (desc != nil))
	if desc != nil {
		beego.Debug("---onononononononononon------ desc = ", *desc)
	}

	rs, e := app.Model().GetReservations(search, userID, localID, confirmed,
		pending, date, notBeforeDate, localAdminID, limit, offset, orderby, desc)

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
	if c.GetAuthor().HaveRol(models.RolSuperadmin) {
		return
	}
	_, e := app.Model().GetLocalAdmin(localID, userID)
	if e == models.ErrNoRows {
		c.WE(fmt.Errorf("Forbbiden"), 403)
	}
	c.WE(e, 500)
}
