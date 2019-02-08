package controllers

import (
	"fmt"

	"github.com/mdiazp/sigel-server/api/app"
	"github.com/mdiazp/sigel-server/api/models"
)

// BaseNotificationController ...
type BaseNotificationController struct {
	BaseController
}

// GetNotification ...
func (c *BaseNotificationController) GetNotification() *models.Notification {
	n := app.Model().NewNotification()
	n.ID = *(c.ReadInt("notification_id", true))
	e := n.Load()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	if n.UserID != c.GetAuthor().ID {
		c.WE(fmt.Errorf("Forbidden"), 403)
	}
	return n
}

// GetNotifications ...
func (c *BaseNotificationController) GetNotifications() *[]*models.Notification {
	limit, offset, orderby, desc := c.ReadPagOrder()
	userID := c.ReadInt("user_id")
	readed := c.ReadBool("readed")
	sdate := c.ReadString("date")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	if orderby == nil {
		tmp := "creation_time"
		orderby = &tmp
		tmp2 := true
		desc = &tmp2
	}

	ns, e := app.Model().GetNotifications(limit, offset, orderby, desc, userID, date, readed)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return ns
}

// GetNotificationsCount ...
func (c *BaseNotificationController) GetNotificationsCount() int {
	userID := c.ReadInt("user_id")
	readed := c.ReadBool("readed")
	sdate := c.ReadString("date")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	cnt, e := app.Model().GetNotificationsCount(userID, date, readed)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return cnt
}

// SetUserNotificationsAsReaded ...
func (c *BaseNotificationController) SetUserNotificationsAsReaded() {
	userID := c.ReadInt("user_id")

	e := app.Model().SetUserNotificationsAsReaded(*userID)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
}

// ReadNotification ...
func (c *BaseNotificationController) ReadNotification() {
	userID := c.ReadInt("user_id")
	nID := c.ReadInt("notification_id")

	n := app.Model().NewNotification()
	n.ID = *nID

	e := n.Load()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	if n.UserID != *userID {
		c.WE(fmt.Errorf("Forbidden"), 403)
	}

	n.Readed = true
	e = n.Update()
	c.WE(e, 500)
}
