package controllers

/*
import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminLocalsController struct {
	LocalsBaseController
}

// @Title Retrieve Local Info
// @Description Get local info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/:id [get]
func (this *AdminLocalsController) Get() {
	obj := this.get()
	this.Data["json"] = obj
	this.ServeJSON()
}

// @Title Create new local
// @Description Create new local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local		body	models.Local	true		"New Local"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [post]
func (this *AdminLocalsController) Post() {
	pthis := &this.Controller

	local := models.Local{}
	ReadInputBody(pthis, &local)

	Validate(pthis, &local)
	_, e := AppModel.GetAreaById(local.AreaId)
	if e == models.ErrResultNotFound {
		wre(pthis, 400)
	}
	if e != nil {
		wre(pthis, 500)
	}

	obj, e := AppModel.CreateLocal(local)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.Data["json"] = obj
	this.ServeJSON()
}

// @Title Update Local
// @Description Edit local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Local id"
// @Param	local		body	models.Local	true		"Edited Local"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/:id [put]
func (this *AdminLocalsController) Put() {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	local := models.Local{}
	ReadInputBody(pthis, &local)

	local.Id = id

	Validate(pthis, &local)
	_, e = AppModel.GetAreaById(local.AreaId)
	if e == models.ErrResultNotFound {
		wre(pthis, 400)
	}
	if e != nil {
		wre(pthis, 500)
	}

	_, e = AppModel.UpdateLocal(local)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.ServeJSON()
}

// @Title Delete Local
// @Description Remove local by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	string	true		"Local id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/:id [delete]
func (this *AdminLocalsController) Delete() {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}
	e = AppModel.DeleteLocal(id)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	this.ServeJSON()
}

// @Title Get Locals List
// @Description Get locals list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	true		"Order Desc (true or false)"
// @Param	area_id		query	int	false		"Order Desc"
// @Param	enable_to_reserve		query	string	false		"Local Property (true o false)"
// @Success 200 {object} []models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (this *AdminLocalsController) List() {
	l := this.list()
	this.Data["json"] = l
	this.ServeJSON()
}
*/
