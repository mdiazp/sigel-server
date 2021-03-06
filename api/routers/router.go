// @APIVersion 1.0.0
// @Title api-sirel
// @Description Documentation for api-sirel
// @Contact di@upr.edu.cu
// @License UPR
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/mdiazp/sigel-server/api/controllers"
	"github.com/mdiazp/sigel-server/api/controllers/admin"
	"github.com/mdiazp/sigel-server/api/controllers/private"
	"github.com/mdiazp/sigel-server/api/controllers/public"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "authHd"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/public",
			beego.NSInclude(
				&public.LoginController{},
				&public.UsersController{},
				&public.LocalsController{},
				&public.AreasController{},
				&public.ReservationsController{},
				&public.InfoController{},
			),
		),
		beego.NSNamespace("/private",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&private.LogoutController{},
				&private.ProfileController{},
				&private.ReservationsController{},
				&private.NotificationsController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSBefore(controllers.AdminFilter),
			beego.NSInclude(
				&admin.UsersController{},
				&admin.AreasController{},
				&admin.LocalsController{},
				&admin.ReservationsController{},
				&admin.NotificationsController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.ErrorController(&controllers.ErrorController{})
}
