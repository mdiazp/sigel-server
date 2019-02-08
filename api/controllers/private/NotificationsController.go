package private

import (
	"strconv"

	"github.com/mdiazp/sigel-server/api/controllers"
)

// NotificationsController ...
type NotificationsController struct {
	controllers.BaseNotificationController
}

// GetNotification ...
// @Title Get Notification
// @Description Get notification by username
// @Param	authHd		header	string	true		"Authentication token"
// @Param	notification_id		query	int	true		"Notification ID"
// @Success 200 {object} model.Notification
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/notification [get]
func (c *NotificationsController) GetNotification() {
	c.Ctx.Input.SetParam("user_id", string(c.GetAuthor().ID))
	c.Data["json"] = c.BaseNotificationController.GetNotification()
	c.ServeJSON()
}

// GetNotifications ...
// @Title Get Notifications
// @Description Get notification by username
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
// @Param	readed		query	string		"(true or false)"
// @Param	date		query	string		"yyyy-mm-dd"
// @Success 200 {object} []model.Notification
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/notifications [get]
func (c *NotificationsController) GetNotifications() {
	c.Ctx.Input.SetParam("user_id", strconv.Itoa(c.GetAuthor().ID))
	c.Data["json"] = c.BaseNotificationController.GetNotifications()
	c.ServeJSON()
}

// GetNotificationsCount ...
// @Title Get Notifications
// @Description Get notification by username
// @Param	authHd		header	string	true		"Authentication token"
// @Param	readed		query	string		"(true or false)"
// @Param	date		query	string		"yyyy-mm-dd"
// @Success 200 int
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/notificationscount [get]
func (c *NotificationsController) GetNotificationsCount() {
	c.Ctx.Input.SetParam("user_id", strconv.Itoa(c.GetAuthor().ID))
	c.Data["json"] = c.BaseNotificationController.GetNotificationsCount()
	c.ServeJSON()
}

// SetUserNotificationsAsReaded ...
// @Title Set all user's notifications as readed
// @Description Set all user's notifications as readed
// @Param	authHd		header	string	true		"Authentication token"
// @Success 200 OK
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/readallnotifications [patch]
func (c *NotificationsController) SetUserNotificationsAsReaded() {
	c.Ctx.Input.SetParam("user_id", strconv.Itoa(c.GetAuthor().ID))
	c.BaseNotificationController.SetUserNotificationsAsReaded()
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// ReadNotification ...
// @Title Read Notification
// @Description Read Notification
// @Param	authHd		header	string	true		"Authentication token"
// @Success 200 OK
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/readnotification [patch]
func (c *NotificationsController) ReadNotification() {
	c.Ctx.Input.SetParam("user_id", strconv.Itoa(c.GetAuthor().ID))
	c.BaseNotificationController.ReadNotification()
	c.Data["json"] = "OK"
	c.ServeJSON()
}
