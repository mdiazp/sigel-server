package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// ErrorController ...
type ErrorController struct {
	beego.Controller
}

// Error400 ...
func (c *ErrorController) Error400() {
	c.he(400)
}

// Error401 ...
func (c *ErrorController) Error401() {
	c.he(401)
}

// Error403 ...
func (c *ErrorController) Error403() {
	c.he(403)
}

// Error404 ...
func (c *ErrorController) Error404() {
	c.he(404)
}

// Error500 ...
func (c *ErrorController) Error500() {
	c.he(500)
}

// Error503 ...
func (c *ErrorController) Error503() {
	c.he(503)
}

// he ...
func (c *ErrorController) he(code int) {
	c.Ctx.Output.SetStatus(404)
	c.Data["json"] = http.StatusText(404)
	c.ServeJSON()
}
