package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) WE(e error, statusCode int, ms ...interface{}) {
	if e == nil {
		return
	}
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
	this.WE(e, 400)
}

func (this *BaseController) GetAuthor() models.User {
	// Author of request must be loggued
	u, e := GetAuthorFromInputData(this.Ctx)
	if e != nil {
		// Then the authenticator filter fail
		beego.Error(e.Error())
		this.WE(e, 500)
	}
	return u
}

func (this *BaseController) Validate(obj interface{}) {
	valid := validation.Validation{}
	ok, e := valid.Valid(obj)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
	if !ok {
		beego.Debug(valid.Errors)
		this.WE(errors.New("bad request"), 400)
	}
}

func (this *BaseController) ReadPagAndOrdOptions() PagAndOrdOptions {
	var (
		opt PagAndOrdOptions
		e   error
	)

	opt.Limit, e = this.GetInt("limit", 20)
	opt.Offset, e = this.GetInt("offset", 0)
	opt.OrderBy = this.GetString("orderby", "id")
	opt.Desc, e = this.GetBool("desc", false)

	this.WE(e, 400)

	return opt
}

func (this *BaseController) Fmtorder(opt *PagAndOrdOptions) string {
	exp := opt.OrderBy
	if opt.Desc {
		exp = "-" + exp
	}
	return exp
}
