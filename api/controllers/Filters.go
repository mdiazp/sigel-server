package controllers

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

// AuthHd ...
const AuthHd = "authHd"

// AuthFilter ...
func AuthFilter(ctx *context.Context) {
	/*
	if ok, _ := beego.AppConfig.Bool("DISABLE_AUTH"); ok {
		u := app.Model().NewUser()
		e := app.Model().RetrieveOne(u, "k_user.username=$1", "manuel.diaz")
		if e != nil {
			wrec(ctx, 500)
		}
		ctx.Input.SetData("Author", u)
		return
	}
	*/

	username, e := app.Crypto().Decrypt(ctx.Input.Header(AuthHd))
	if e != nil {
		/*beego.Debug(e)*/
		wrec(ctx, 401)
		return
	}

	u := app.Model().NewUser()
	if username != "SIREL" {
		ku := app.Model().NewUser()
		e = app.Model().RetrieveOne(ku, "username=$1", username)

		if e != nil {
			if e == models.ErrNoRows {
				wrec(ctx, 401)
				return
			}
			beego.Error(e.Error())
			wrec(ctx, 500)
			return
		}

		if !ku.Enable {
			wrec(ctx, 401)
			return
		}

		u = ku
	} else {
		u.Username = "SIREL"
		u.Rol = models.RolSuperadmin
	}

	ctx.Input.SetData("Author", u)
}

// AdminFilter ...
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
