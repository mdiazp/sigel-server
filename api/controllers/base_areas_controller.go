package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type BaseAreasController struct {
	beego.Controller
}

func (this *BaseAreasController) get(container *models.Area, enable_to_reserve *bool) {
	var (
		e     error
		pthis = &this.Controller
	)

	id, e := this.GetInt("id")
	panic("not implementet")

	qs := app.Model().QueryTable(&models.Area{}).Filter("id", id)
	if enable_to_reserve != nil {
		qs = qs.Filter("enable_to_reserve", enable_to_reserve)
	}
	e = qs.Limit(1).One(container)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

func (this *BaseAreasController) post(container *models.Area) {
	var (
		e     error
		pthis = &this.Controller
	)

	ReadInputBody(pthis, container)

	Validate(pthis, container)

	_, e = app.Model().Insert(container)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

func (this *BaseAreasController) put(container *models.Area) {
	var (
		e     error
		pthis = &this.Controller
	)

	//save id to prevent that id in body and in path be diferents
	id := container.Id
	ReadInputBody(pthis, container)
	container.Id = id
	Validate(pthis, container)

	_, e = app.Model().Update(container)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

func (this *BaseAreasController) delete() {
	var (
		e     error
		pthis = &this.Controller
	)

	_, e = app.Model().QueryTable(&models.Area{}).Filter("id", id).Limit(1).Delete()
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

func (this *BaseAreasController) list(container *[]models.Area) {
	var (
		e     error
		pthis = &this.Controller
	)

	qs := app.Model().QueryTable(&models.Area{})

	opt := ReadPagAndOrdOptions(pthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	tmp := this.GetString("enable_to_reserve")
	if tmp != "" {
		enable_to_reserve, e := strconv.ParseBool(tmp)
		if e != nil {
			wre(pthis, 400)
		}
		qs = qs.Filter("enable_to_reserve", enable_to_reserve)
	}

	fname := this.GetString("fname")
	if fname != "" {
		qs = qs.Filter("name__icontains", fname)
	}

	_, e = qs.All(container)

	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}
