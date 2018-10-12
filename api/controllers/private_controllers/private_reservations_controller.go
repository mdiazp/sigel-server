package private_controllers

import (
	"github.com/mdiazp/sirel-server/api/controllers"
)

type PrivateReservationsController struct {
	controllers.BaseReservationsController
}

// @Title Post Reservation
// @Description Autor make new reservation
// @Param	authHd		header	string	true		"Authentication token"
// @Param	reservation		body	models.Reservation	true		"New Reservation"
// @Success 200 {object} models.Reservation
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservation [post]
func (this *PrivateReservationsController) Post() {

}
