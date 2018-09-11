package public_controllers

import (
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type PublicAreasController struct {
	controllers.BaseAreasController
}

// @Title Retrieve public area info
// @Description Get area info by id
// @Param	id		query	int	true		"Area id"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (this *PublicAreasController) Get() {
	id, e := this.GetInt("id")
	this.WE(e, 400)

	o := models.Area{}
	enable_to_reserve := true

	this.BaseAreasController.Show(id, &o, &enable_to_reserve)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Get public areas list
// @Description Get public areas list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Param	fname		query	string	false		"Search in name"
// @Success 200 {object} []models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *PublicAreasController) List() {
	var l []models.Area
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	this.BaseAreasController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}
