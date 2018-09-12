package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type BaseAreasController struct {
	BaseController
}

func (this *BaseAreasController) Show(container *models.Area) {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	qs := app.Model().QueryTable(&models.Area{}).Filter("id", id)
	enable_to_reserve := this.GetString("enable_to_reserve")
	if enable_to_reserve != "" {
		qs = qs.Filter("enable_to_reserve", enable_to_reserve)
	}
	e = qs.Limit(1).One(container)
	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) Create(container *models.Area) {
	var (
		e error
	)

	this.ReadInputBody(container)

	this.Validate(container)

	_, e = app.Model().Insert(container)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) Update(container *models.Area) {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	//save id to prevent that id in body and in path be diferents
	this.ReadInputBody(container)
	container.Id = id
	this.Validate(container)

	_, e = app.Model().Update(container)
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) Remove() {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	_, e = app.Model().QueryTable(&models.Area{}).Filter("id", id).Limit(1).Delete()
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) List(container *[]models.Area) {
	var (
		e error
	)

	qs := app.Model().QueryTable(&models.Area{})

	opt := this.ReadPagAndOrdOptions()
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(this.Fmtorder(&opt))
	}

	tmp := this.GetString("enable_to_reserve")
	if tmp != "" {
		enable_to_reserve, e := strconv.ParseBool(tmp)
		this.WE(e, 400)
		qs = qs.Filter("enable_to_reserve", enable_to_reserve)
	}

	fname := this.GetString("fname")
	if fname != "" {
		qs = qs.Filter("name__icontains", fname)
	}

	_, e = qs.All(container)

	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}
