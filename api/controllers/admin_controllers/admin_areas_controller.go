package admin_controllers

import (
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminAreasController struct {
	controllers.BaseAreasController
}

// @Title Retrieve Area Info
// @Description Get area info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Area id"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (this *AdminAreasController) Get() {
	o := models.Area{}

	this.BaseAreasController.Show(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Create new area
// @Description Create new area (role admin required)
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
func (this *AdminAreasController) Post() {
	o := models.Area{}

	this.BaseAreasController.Create(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Update Area
// @Description Edit area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Area id"
// @Param	area		body	models.Area	true		"Edited Area"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [put]
func (this *AdminAreasController) Put() {
	o := models.Area{}

	this.BaseAreasController.Update(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Delete Area
// @Description Remove area by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	string	true		"Area id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [delete]
func (this *AdminAreasController) Remove() {
	this.BaseAreasController.Remove()
	this.ServeJSON()
}

// @Title Get Areas List
// @Description Get areas list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Param	enable_to_reserve		query	string	false		"Area Property (true o false)"
// @Param	fname		query	string	false		"Search in name"
// @Success 200 {object} []models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *AdminAreasController) List() {
	var l []models.Area
	this.BaseAreasController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}
