package filters

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

const AuthHd = "authHd"

func Auth(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/api/login") {
		return
	}

	if ok, _ := beego.AppConfig.Bool("DISABLE_AUTH"); ok {
		u := models.User{
			Username: "manuel.diaz",
			Rol:      "Admin",
		}
		ctx.Input.SetData("Author", u)
		return
	}

	username, e := controllers.GetCrypto().Decrypt(ctx.Input.Header(AuthHd))
	if e != nil {
		beego.Debug(e)
		wre(ctx, 401)
		return
	}

	var u models.User
	if username != "SIREL" {
		u, e = controllers.AppModel.GetUserByUsername(username)
		if e != nil {
			if e == models.ErrResultNotFound {
				wre(ctx, 401)
				return
			}
			beego.Error(e.Error())
			wre(ctx, 500)
			return
		}

		if !u.Enable {
			wre(ctx, 401)
			return
		}
	} else {
		u = models.User{
			Username: "SIREL",
			Rol:      models.RolSuperadmin,
		}
	}

	ctx.Input.SetData("Author", u)
}

func Admin(ctx *context.Context) {
	if !strings.HasPrefix(ctx.Input.URL(), "/admin") {
		return
	}
	u, e := controllers.GetAuthorFromInputData(ctx)
	if e != nil {
		beego.Error(e.Error())
		wre(ctx, 500)
		return
	}
	if !u.HaveRol("Admin") {
		wre(ctx, 403)
		return
	}
}

func wre(ctx *context.Context, statusCode int, ms ...interface{}) {
	ctx.Output.SetStatus(statusCode)
	if len(ms) > 0 {
		ctx.Output.JSON(ms, false, false)
	} else {
		ctx.Output.JSON(http.StatusText(statusCode), false, false)
	}
}
