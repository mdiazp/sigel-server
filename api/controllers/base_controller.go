package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) WE(e error, statusCode int, ms ...interface{}) {
	this.Ctx.Output.SetStatus(statusCode)
	if len(ms) > 0 {
		this.Data["json"] = ms[0]
	} else {
		this.Data["json"] = http.StatusText(statusCode)
	}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ReadInputBody(obj interface{}) {
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &obj)
	if e != nil {
		wre(this, 400)
	}
}

func (this *BaseController) GetAuthor() models.User {
	// Author of request must be loggued
	u, e := GetAuthorFromInputData(this.Ctx)
	if e != nil {
		// Then the authenticator filter fail
		beego.Error(e.Error())
		wre(this, 500)
	}
	return u
}

func (this *BaseController) Validate(obj interface{}) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		beego.Error(err.Error())
		wre(this, 500)
	}
	if !ok {
		beego.Debug(valid.Errors)
		wre(this, 400)
	}
}

func ReadPagAndOrdOptions(this *beego.Controller) PagAndOrdOptions {
	var (
		opt PagAndOrdOptions
		e   error
	)

	opt.Limit, e = this.GetInt("limit", 20)
	opt.Offset, e = this.GetInt("offset", 0)
	opt.OrderBy = this.GetString("orderby", "id")
	opt.Desc, e = this.GetBool("desc", false)

	if e != nil {
		wre(this, 400)
	}

	return opt
}

func fmtorder(opt *PagAndOrdOptions) string {
	exp := opt.OrderBy
	if opt.Desc {
		exp = "-" + exp
	}
	return exp
}

/*
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type baseController struct {
	beego.Controller
}

// container must be a pointer
func (this *baseController) get(container models.Area) {
	var (
		e     error
		pthis = &this.Controller
	)
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}



	e = app.Model().QueryTable(&models.Area{}).Filter("id", id).Limit(1).One(container)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

// container must be a pointer
func (this *baseController) post(container interface{}, readBody bool) {
	var (
		e     error
		pthis = &this.Controller
	)

	if readBody {
		ReadInputBody(pthis, container)
	}
	Validate(pthis, container)

	_, e = app.Model().Insert(container)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

// container must be a pointer
func (this *baseController) put(container models.Area, readBody bool) {

	var (
		e     error
		pthis = &this.Controller
	)
	id, e = this.GetInt("id")
	if e != nil {
			wre(pthis, 400)
		}
	}

	if readBody {
		ReadInputBody(pthis, container)
	}
	(*setId)(*id, container)
	Validate(pthis, container)

	_, e = app.Model().Update(container)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

//o must be a pointer
func (this *baseController) delete(o interface{}, setId *func(int, interface{})) {
	var (
		e     error
		pthis = &this.Controller
	)
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}
	(*setId)(id, o)
	_, e = app.Model().Delete(o)
	if e == models.ErrResultNotFound {
		wre(pthis, 404)
	}
	this.Data["json"] = "OK"
	this.ServeJSON()
}

// containerList must be *[]interface{}
func (this *baseController) list(qs orm.QuerySeter, containerList interface{}, fp ...filter_param) {
	var (
		e     error
		pthis = &this.Controller
	)

	opt := ReadPagAndOrdOptions(pthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)

	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	_, e = qs.All(containerList)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
}

type filter_param struct {
	param_name    string
	default_value *string
	valid_options []string
}
*/
