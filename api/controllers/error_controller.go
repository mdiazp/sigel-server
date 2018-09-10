package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error400() {
	this.he(400)
}

func (this *ErrorController) Error401() {
	this.he(401)
}

func (this *ErrorController) Error403() {
	this.he(403)
}

func (this *ErrorController) Error404() {
	this.he(404)
}

func (this *ErrorController) Error500() {
	this.he(500)
}

func (this *ErrorController) Error503() {
	this.he(503)
}

func (this *ErrorController) he(code int) {
	this.Ctx.Output.SetStatus(404)
	this.Data["json"] = http.StatusText(404)
	this.ServeJSON()
}
