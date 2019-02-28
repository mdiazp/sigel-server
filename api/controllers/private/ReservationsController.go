package private

import (
	"fmt"
	"strconv"

	"github.com/mdiazp/sigel-server/api/app"
	"github.com/mdiazp/sigel-server/api/controllers"
	"github.com/mdiazp/sigel-server/api/models"
)

// ReservationsController ...
type ReservationsController struct {
	controllers.BaseReservationsController
}

// Post ...
// @Title Create new reservation
// @Description Create new reservation (role user required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservation		body	controllers.ReservationToCreate	true		"New Reservation"
// @Success 200 {object} models.ReservationInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/reservation [post]
func (c *ReservationsController) Post() {
	c.AccessControl(models.RolUser)
	c.Data["json"] = c.BaseReservationsController.Create()
	c.ServeJSON()
}

// Confirm ...
// @Title Confirm reservation
// @Description Confirm reservation (role user required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservationID		query	int	true		"Reservation ID"
// @Success 200 {object} models.ReservationInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/reservation [patch]
func (c *ReservationsController) Confirm() {
	c.AccessControl(models.RolUser)
	c.Data["json"] = c.BaseReservationsController.Confirm()
	c.ServeJSON()
}

// Cancel ...
// @Title Cancel reservation
// @Description Cancel reservation (role user required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservationID		query	int	true		"Reservation ID"
// @Success 200 string
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/reservation [delete]
func (c *ReservationsController) Cancel() {
	r := c.LoadReservation()
	u := c.GetAuthor()

	/*beego.Debug("u.ID = ", u.ID, " r.UserID = ", r.UserID, " r.Pending = ", r.Pending)*/

	if u.ID != r.UserID || !r.Pending {
		c.WE(fmt.Errorf("Invalid operation"), 403)
	}
	e := app.Model().Delete(r)
	c.WE(e, 500)

	c.Data["json"] = "OK"
	c.ServeJSON()
}

// List ...
// @Title Get user's reservations list
// @Description Get user's reservations list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
// @Param	user_id		query	int	false		"User ID"
// @Param	local_id		query	int	false		"Local ID"
// @Param	confirmed		query	string	false		"true or false"
// @Param	pending		query	string	false		"true or false"
// @Param	date		query	string		"yyyy-mm-dd"
// @Param	not_before_date		query	string		"yyyy-mm-dd"
// @Param	search		query	string	false		"Search in activity name"
// @Success 200 {object} []models.ReservationInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/reservations [get]
func (c *ReservationsController) List() {
	c.Ctx.Input.SetParam("user_id", strconv.Itoa(c.GetAuthor().ID))
	c.Ctx.Input.SetParam("localAdminID", "")

	c.Data["json"] = c.BaseReservationsController.List().Reservations
	c.ServeJSON()
}
