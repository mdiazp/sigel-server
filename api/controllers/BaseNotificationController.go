package controllers

import (
	"fmt"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
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
	sdate := c.ReadString("date")
	date, e := app.Model().NewDate(sdate)
	if e != nil {
		c.WE(e, 400)
	}

	ns, e := app.Model().GetNotifications(limit, offset, orderby, desc, userID, date)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return ns
}
