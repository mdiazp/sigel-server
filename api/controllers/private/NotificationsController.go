package private

import "github.com/mdiazp/sirel-server/api/controllers"

// NotificationController ...
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
// @router /profile/notification [get]
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
// @Param	orderDesc		query	string	false		"true or false"
// @Param	date		query	string		"yyyy-mm-dd"
// @Success 200 {object} []model.Notification
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /profile/notification [get]
func (c *NotificationsController) GetNotifications() {
	c.Ctx.Input.SetParam("user_id", string(c.GetAuthor().ID))
	c.Data["json"] = c.BaseNotificationController.GetNotifications()
	c.ServeJSON()
}
