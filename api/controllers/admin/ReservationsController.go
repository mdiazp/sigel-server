package admin

import (
	"strconv"

	"github.com/mdiazp/sigel-server/api/controllers"
	"github.com/mdiazp/sigel-server/api/models"
)

// ReservationsController ...
type ReservationsController struct {
	controllers.BaseReservationsController
}

// Get ...
// @Title Retrieve reservation info
// @Description Get reservation's info by id
// @Param	authHd		header	string	true		"Authentication token"
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

// Accept ...
// @Title Accept reservation
// @Description Accept reservation, pending will be false
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservation_id		query	int	true		"Reservation id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservation [patch]
func (c *ReservationsController) Accept() {
	c.BaseReservationsController.AcceptReservation()
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// Refuse ...
// @Title Refuse reservation
// @Description Refuse reservation and delete it
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservation_id		query	int	true		"Reservation id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservation [delete]
func (c *ReservationsController) Refuse() {
	c.BaseReservationsController.RefuseReservation()
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// List ...
// @Title Get public reservation list
// @Description Get public areas list
// @Param	authHd		header	string	true		"Authentication token"
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
	if !c.GetAuthor().HaveRol(models.RolSuperadmin) {
		c.Ctx.Input.SetParam("localAdminID", strconv.Itoa(c.GetAuthor().ID))
	}
	rs := c.BaseReservationsController.List().Reservations

	c.Data["json"] = rs
	c.ServeJSON()
}

// List2 ...
// @Title Get public reservation list
// @Description Get public areas list
// @Param	authHd		header	string	true		"Authentication token"
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
// @Success 200 {object} []controllers.ReservationWithusername
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservations2 [get]
func (c *ReservationsController) List2() {
	if !c.GetAuthor().HaveRol(models.RolSuperadmin) {
		c.Ctx.Input.SetParam("localAdminID", strconv.Itoa(c.GetAuthor().ID))
	}
	rs := c.BaseReservationsController.List2()

	c.Data["json"] = rs
	c.ServeJSON()
}

// ReservationsCount ...
// @Title Get public reservations count
// @Description Get public reservations count
// @Param	authHd		header	string	true		"Authentication token"
// @Param	user_id		query	int	false		"User ID"
// @Param	local_id		query	int	false		"Local ID"
// @Param	confirmed		query	string	false		"true or false"
// @Param	pending		query	string	false		"true or false"
// @Param	date		query	string		"yyyy-mm-dd"
// @Param	search		query	string	false		"Search in activity name"
// @Success 200 int
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservationscount [get]
func (c *ReservationsController) ReservationsCount() {
	if !c.GetAuthor().HaveRol(models.RolSuperadmin) {
		c.Ctx.Input.SetParam("localAdminID", strconv.Itoa(c.GetAuthor().ID))
	}
	c.Data["json"] = c.BaseReservationsController.Count()
	c.ServeJSON()
}
