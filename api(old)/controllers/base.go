package controllers

import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type Identity interface {
	SetId(int)
}

type BaseController struct {
	beego.Controller
	m Model
}

func (this *BaseController) get() Identity {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	o, e := this.m.GetById(id)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	return o
}

func (this *BaseController) post() Identity {
	pthis := &this.Controller

	o := this.m.GetZero()
	ReadInputBody(pthis, &o)

	Validate(pthis, &o)

	o, e := this.m.Create(o)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.Data["json"] = o
}

func (this *BaseController) put() Identity {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	o := this.m.GetZero()
	ReadInputBody(pthis, &o)

	o.SetId(id)
	Validate(pthis, &o)

	_, e = this.m.Update(o)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.ServeJSON()
}

func (this *BaseController) Delete() {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}
	e = this.m.Delete(id)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	this.ServeJSON()
}

func (this *BaseController) list(containerList interface{}) {
	pthis := &this.Controller
	qs := this.m.GetQuerySeter()

	opt := ReadPagAndOrdOptions(pthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)

	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	_, e := qs.All(containerList)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}
