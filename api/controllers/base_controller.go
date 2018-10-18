package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/mdiazp/sirel-server/api/models/models2"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) WE(e error, statusCode int, ms ...interface{}) {
	if e == nil {
		return
	}

	if statusCode == 500 {
		beego.Debug(e.Error())
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

func (this *BaseController) GetAuthor() *models2.User {
	// Author of request must be loggued
	u, e := GetAuthorFromInputData(this.Ctx)
	this.WE(e, 500)
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
		beego.Debug(fmt.Sprint(obj))
		beego.Debug(valid.Errors)
		this.WE(errors.New("bad request"), 400)
	}
}

func (this *BaseController) ReadPagAndOrdOptions(defaultOrderByOption string, orderByOptions ...string) PagAndOrdOptions {
	var (
		opt PagAndOrdOptions
		e   error
		ok  bool
	)

	opt.Limit, e = this.GetInt("limit", 20)
	opt.Offset, e = this.GetInt("offset", 0)
	opt.OrderBy = this.GetString("orderby", "id")

	ok = false

	for _, o := range orderByOptions {
		if o == opt.OrderBy {
			ok = true
			break
		}
	}

	if !ok {
		opt.OrderBy = defaultOrderByOption
	}

	opt.orderDirection = this.GetString("orderDirection", "asc")
	if opt.orderDirection != "asc" && opt.orderDirection != "desc" {
		e = errors.New(fmt.Sprintf("orderDirection have an invalid value: %s", opt.orderDirection))
		beego.Debug(e.Error())
	}

	this.WE(e, 400)

	return opt
}

func (this *BaseController) Fmtorder(opt *PagAndOrdOptions) string {
	exp := opt.OrderBy
	if opt.orderDirection == "desc" {
		exp = "-" + exp
	}
	return exp
}

// ReadString ...
func (this *BaseController) ReadString(name string, required ...bool) *string {
	tmp := this.GetString(name)
	if tmp != "" {
		return &tmp
	} else if len(required) > 0 && required[0] {
		this.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}

// ReadInt ...
func (this *BaseController) ReadInt(name string, required ...bool) *int {
	tmp := this.GetString(name)
	if tmp != "" {
		c, e := strconv.Atoi(tmp)
		this.WE(e, 400)
		return &c
	} else if len(required) > 0 && required[0] {
		this.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}

// ReadBool ...
func (this *BaseController) ReadBool(name string, required ...bool) *bool {
	tmp := this.GetString(name)
	if tmp != "" {
		c, e := strconv.ParseBool(tmp)
		this.WE(e, 400)
		return &c
	} else if len(required) > 0 && required[0] {
		this.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}
