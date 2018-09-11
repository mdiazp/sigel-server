package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type BaseLocalsController struct {
	BaseController
}

func (this *BaseLocalsController) Show(id int, container *models.Local, enable_to_reserve *bool) {
	var (
		e error
	)

	qs := app.Model().QueryTable(&models.Local{}).Filter("id", id)
	if enable_to_reserve != nil {
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

func (this *BaseLocalsController) Create(container *models.Local) {
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

func (this *BaseLocalsController) Update(id int, container *models.Local) {
	var (
		e error
	)

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

func (this *BaseLocalsController) Remove(id int) {
	var (
		e error
	)

	_, e = app.Model().QueryTable(&models.Local{}).Filter("id", id).Limit(1).Delete()
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseLocalsController) List(container *[]models.Local) {
	var (
		e error
	)

	qs := app.Model().QueryTable(&models.Local{})

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

	area_id, e := this.GetInt("area_id")
	if area_id > 0 {
		qs = qs.Filter("area_id", area_id)
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
