package controllers

/*
import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminAreasController struct {
	AreasBaseController
}

// @Title Retrieve Area Info
// @Description Get area info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Area id"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/:id [get]
func (this *AdminAreasController) Get() {
	a := this.get()
	this.Data["json"] = a
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

// @Title Update Area
// @Description Edit area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Area id"
// @Param	area		body	models.Area	true		"Edited Area"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/:id [put]
func (this *AdminAreasController) Put() {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	area := models.Area{}
	ReadInputBody(pthis, &area)

	area.Id = id

	Validate(pthis, &area)

	_, e = AppModel.UpdateArea(area)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.ServeJSON()
}

// @Title Delete Area
// @Description Remove area by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	string	true		"Area id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/:id [delete]
func (this *AdminAreasController) Delete() {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}
	e = AppModel.DeleteArea(id)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	this.ServeJSON()
}

// @Title Get Areas List
// @Description Get areas list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	true		"Order Desc"
// @Param	enable_to_reserve		query	string	false		"Area Property (true o false)"
// @Success 200 {object} []models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *AdminAreasController) List() {
	areas := this.list()
	this.Data["json"] = areas
	this.ServeJSON()
}
*/
