package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/mdiazp/sirel-server/api/models"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// WE ...
func (c *BaseController) WE(e error, statusCode int, ms ...interface{}) {
	if e == nil {
		return
	}

	if statusCode == 400 {
		beego.Debug(e.Error())
	}

	if statusCode == 500 {
		beego.Debug(e.Error())
	}

	if statusCode == 401 {
		beego.Debug(e.Error())
	}

	c.Ctx.Output.SetStatus(statusCode)
	if len(ms) > 0 {
		c.Data["json"] = ms[0]
	} else {
		c.Data["json"] = http.StatusText(statusCode)
	}
	c.ServeJSON()
	c.StopRun()
}

// ReadInputBody ...
func (c *BaseController) ReadInputBody(obj interface{}) {
	e := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	c.WE(e, 400)
}

// GetAuthor ...
func (c *BaseController) GetAuthor() *models.User {
	// Author of request must be loggued
	u, e := GetAuthorFromInputData(c.Ctx)
	c.WE(e, 500)
	return u
}

// Validate ...
func (c *BaseController) Validate(obj interface{}) {
	valid := validation.Validation{}
	ok, e := valid.Valid(obj)
	if e != nil {
		beego.Error(e.Error())
		c.WE(e, 500)
	}
	if !ok {
		beego.Debug(fmt.Sprint(obj))
		beego.Debug(valid.Errors)
		c.WE(errors.New("bad request"), 400)
	}
}

// ReadPagOrder ...
func (c *BaseController) ReadPagOrder() (*int, *int, *string, *bool) {
	limit := c.ReadInt("limit")
	offset := c.ReadInt("offset")
	orderby := c.ReadString("orderby")
	desc := c.ReadBool("desc")
	return limit, offset, orderby, desc
}

// Fmtorder ...
func (c *BaseController) Fmtorder(opt *PagAndOrdOptions) string {
	exp := opt.OrderBy
	if opt.orderDirection == "desc" {
		exp = "-" + exp
	}
	return exp
}

// ReadString ...
func (c *BaseController) ReadString(name string, required ...bool) *string {
	tmp := c.GetString(name)
	if tmp != "" {
		return &tmp
	} else if len(required) > 0 && required[0] {
		c.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}

// ReadInt ...
func (c *BaseController) ReadInt(name string, required ...bool) *int {
	tmp := c.GetString(name)
	beego.Debug(name + " = -" + tmp + "-")
	if tmp != "" {
		x, e := strconv.Atoi(tmp)
		c.WE(e, 400)
		return &x
	} else if len(required) > 0 && required[0] {
		c.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}

// ReadBool ...
func (c *BaseController) ReadBool(name string, required ...bool) *bool {
	tmp := c.GetString(name)
	if tmp != "" {
		x, e := strconv.ParseBool(tmp)
		c.WE(e, 400)
		return &x
	} else if len(required) > 0 && required[0] {
		c.WE(fmt.Errorf("%s is missing in the input", name), 400)
	}
	return nil
}

// ReadObjectInBody ...
func (c *BaseController) ReadObjectInBody(name string, o interface{}, required ...bool) bool {
	if len(c.Ctx.Input.RequestBody) == 0 {
		if len(required) > 0 && required[0] {
			c.WE(fmt.Errorf("Empty %s in body", name), 400)
		}
		return false
	}
	e := json.Unmarshal(c.Ctx.Input.RequestBody, o)
	c.WE(e, 400)
	return true
}

// AccessControl ...
func (c *BaseController) AccessControl(rol string) {
	author := c.GetAuthor()
	if author.HaveRol(rol) {
		return
	}
	c.WE(fmt.Errorf("Forbidden"), 403)
}
