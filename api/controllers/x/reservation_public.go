package controllers

/*
import (
	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

// @Title Create new area
// @Description Create new reservation (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area		body	models.Area	true		"New Area"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [post]
func (this *AdminReservationController) Post() {
	pthis := &this.Controller

	area := models.Area{}
	ReadInputBody(pthis, &area)

	Validate(pthis, &area)

	a, e := AppModel.CreateArea(area)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.Data["json"] = a
	this.ServeJSON()
}
*/
