package controllers

import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminLocalsController struct {
	beego.Controller
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
	var (
		e     error
		pthis = &this.Controller
	)
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	o := models.Local{}
	e = app.Model().QueryTable(&models.Local{}).Filter("id", id).Limit(1).One(&o)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	this.Data["json"] = o
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

	o := models.Local{}

	ReadInputBody(pthis, &o)

	Validate(pthis, &o)

	_, e := app.Model().Insert(&o)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Update Local
// @Description Edit local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Local id"
// @Param	local		body	models.Local	true		"Edited Local"
// @Success 200 {object} models.Local
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

	o := models.Local{}
	ReadInputBody(pthis, &o)

	o.Id = id

	Validate(pthis, &o)

	_, e = app.Model().Update(&o)
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
	o := models.Local{}
	e = app.Model().QueryTable(&models.Local{}).Filter("id", id).Limit(1).One(&o)
	_, e = app.Model().Delete(&o)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	this.ServeJSON()
}

// @Title Get locals List
// @Description Get locals list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	true		"Order Desc"
// @Param	area_id		query	string	false		"Local Property (true o false)"
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
	var (
		e     error
		pthis = &this.Controller
	)

	qs := app.Model().QueryTable(&models.Local{})

	opt := ReadPagAndOrdOptions(pthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	enable_to_reserve, e := this.GetBool("enable_to_reserve")
	if e != nil {
		wre(pthis, 400)
	}
	qs = qs.Filter("enable_to_reserve", enable_to_reserve)

	area_id, e := this.GetInt("area_id")
	if e != nil {
		wre(pthis, 400)
	}
	qs = qs.Filter("area_id", area_id)

	var l []models.Area
	_, e = qs.All(&l)

	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	this.Data["json"] = l
	this.ServeJSON()
}
