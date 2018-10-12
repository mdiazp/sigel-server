package public_controllers

import (
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type PublicReservationsController struct {
	controllers.BaseReservationsController
}

// @Title Retrieve public reservation info
// @Description Get reservation info by id
// @Param	id		query	int	true		"Reservation id"
// @Success 200 {object} models.Reservation
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservation [get]
func (this *PublicReservationsController) Get() {
	o := models.Reservation{}

	this.Show(&o)

	this.Data["json"] = o
	this.ServeJSON()
}