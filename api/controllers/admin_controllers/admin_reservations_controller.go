package admin_controllers

import(
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type AdminReservationsController struct {
	controllers.BaseReservationsController
}

// @Title Get Reservations List
// @Description Get reservations list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	sortorder		query	string	false		"asc or desc"
// @Param	local_id		query	int	false		"Reservation Property"
// @Param	user_id		query	int	false		"Reservation Property"
// @Param	search_in_activity_name		query	string	false		"Search in activity name"
// @Success 200 {object} []models.Reservation
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /reservations [get]
func (this *AdminReservationsController) List() {
	var l []models.Reservation
	this.BaseReservationsController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}