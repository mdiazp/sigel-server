package public_controllers

import (
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type PublicLocalsController struct {
	controllers.BaseLocalsController
}

// @Title Retrieve public local info
// @Description Get local info by id
// @Param	id		query	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (this *PublicLocalsController) Get() {
	id, e := this.GetInt("id")
	this.WE(e, 400)

	o := models.Local{}
	enable_to_reserve := true

	this.BaseLocalsController.Show(id, &o, &enable_to_reserve)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Get public locals list
// @Description Get public locals list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Param	area_id		query	int	false		"Property Local"
// @Param	fname		query	string	false		"Search in name"
// @Success 200 {object} []models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (this *PublicLocalsController) List() {
	var l []models.Local
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	this.BaseLocalsController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}
