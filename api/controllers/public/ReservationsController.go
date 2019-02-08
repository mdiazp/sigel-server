package public

import (
	"github.com/mdiazp/sigel-server/api/controllers"
)

// ReservationsController ...
type ReservationsController struct {
	controllers.BaseReservationsController
}

// Get ...
// @Title Retrieve public area info
// @Description Get reservation's info by id
// @Param	reservation_id		query	int	true		"Reservation id"
// @Success 200 {object} models.ReservationInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservation [get]
func (c *ReservationsController) Get() {
	c.Data["json"] = c.BaseReservationsController.Show()
	c.ServeJSON()
}

// List ...
// @Title Get public reservation list
// @Description Get public areas list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
// @Param	user_id		query	int	false		"User ID"
// @Param	local_id		query	int	false		"Local ID"
// @Param	confirmed		query	string	false		"true or false"
// @Param	pending		query	string	false		"true or false"
// @Param	date		query	string		"yyyy-mm-dd"
// @Param	search		query	string	false		"Search in activity name"
// @Success 200 {object} []models.ReservationInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservations [get]
func (c *ReservationsController) List() {
	c.Ctx.Input.SetParam("localAdminID", "")
	c.Data["json"] = c.BaseReservationsController.List().Reservations
	c.ServeJSON()
}
