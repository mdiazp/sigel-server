package controllers

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

const AuthHd = "authHd"

func AuthFilter(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/api/login") {
		return
	}

	if ok, _ := beego.AppConfig.Bool("DISABLE_AUTH"); ok {
		u := models.KUser{
			Username: "manuel.diaz",
			Rol:      "Admin",
		}
		ctx.Input.SetData("Author", u)
		return
	}

	username, e := app.Crypto().Decrypt(ctx.Input.Header(AuthHd))
	if e != nil {
		beego.Debug(e)
		wrec(ctx, 401)
		return
	}

	var u models.KUser
	if username != "SIREL" {
		e = app.Model().QueryTable(&models.KUser{}).Filter("username", username).Limit(1).One(&u)
		if e != nil {
			if e == models.ErrResultNotFound {
				wrec(ctx, 401)
				return
			}
			beego.Error(e.Error())
			wrec(ctx, 500)
			return
		}

		if !u.Enable {
			wrec(ctx, 401)
			return
		}
	} else {
		u = models.KUser{
			Username: "SIREL",
			Rol:      models.RolSuperadmin,
		}
	}

	ctx.Input.SetData("Author", u)
}

func AdminFilter(ctx *context.Context) {
	if !strings.HasPrefix(ctx.Input.URL(), "/admin") {
		return
	}
	u, e := GetAuthorFromInputData(ctx)
	if e != nil {
		beego.Error(e.Error())
		wrec(ctx, 500)
		return
	}
	if !u.HaveRol("Admin") {
		wrec(ctx, 403)
		return
	}
}

func wrec(ctx *context.Context, statusCode int, ms ...interface{}) {
	ctx.Output.SetStatus(statusCode)
	if len(ms) > 0 {
		ctx.Output.JSON(ms, false, false)
	} else {
		ctx.Output.JSON(http.StatusText(statusCode), false, false)
	}
}
