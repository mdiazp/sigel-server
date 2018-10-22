package private

import (
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
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
// @router /reservation [post]
func (c *ReservationsController) Post() {
	c.AccessControl(models.RolSuperadmin)
	c.Data["json"] = c.BaseReservationsController.Create()
	c.ServeJSON()
}