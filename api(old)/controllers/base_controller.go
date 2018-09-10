package controllers

import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type baseController struct {
	beego.Controller
}

func (this *baseController) get(c interface{}) {
	var (
		e     error
		bthis = &this.Controller
	)

	id, e := this.GetInt("id")
	if e != nil {
		wre(bthis, 400)
	}

	e = models.Model.
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(bthis, 404)
		}
		beego.Error(e.Error())
		wre(bthis, 500)
	}
	return o
}

func (this *baseController) post() models.Identity {
	var (
		e     error
		bthis = &this.Controller
	)

	o := this.z()
	ReadInputBody(bthis, o)

	Validate(bthis, o)
	e = this.ids().Create(o)

	if e != nil {
		beego.Error(e.Error())
		wre(bthis, 500)
	}
	return o
}

func (this *baseController) put() models.Identity {
	var (
		e     error
		bthis = &this.Controller
	)
	id, e := this.GetInt("id")
	if e != nil {
		wre(bthis, 400)
	}

	o := this.z()
	ReadInputBody(bthis, o)

	o.SetId(id)
	Validate(bthis, o)

	e = this.ids().Update(id, o)
	if e != nil {
		beego.Error(e.Error())
		wre(bthis, 500)
	}
	return o
}

func (this *baseController) Delete() {
	var (
		e     error
		bthis = &this.Controller
	)
	id, e := this.GetInt("id")
	if e != nil {
		wre(bthis, 400)
	}
	e = this.ids().Delete(id)
	if e == models.ErrResultNotFound {
		wre(bthis, 404)
	}
	this.ServeJSON()
}

func (this *baseController) List(containerList interface{}) {
	var (
		e     error
		bthis = &this.Controller
	)
	qs := this.ids().GetQuerySeter()

	opt := ReadPagAndOrdOptions(bthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)

	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	_, e = qs.All(containerList)

	if e != nil {
		beego.Error(e.Error())
		wre(bthis, 500)
	}
}
