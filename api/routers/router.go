// @APIVersion 1.0.0
// @Title api-sirel
// @Description Documentation for api-sirel
// @Contact di@upr.edu.cu
// @License UPR
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/controllers/admin_controllers"
	"github.com/mdiazp/sirel-server/api/controllers/private_controllers"
	"github.com/mdiazp/sirel-server/api/controllers/public_controllers"
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
				&public_controllers.LoginController{},
				&public_controllers.PublicAreasController{},
			),
		),
		beego.NSNamespace("/private",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&private_controllers.LogoutController{},
				&private_controllers.ProfileController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSBefore(controllers.AdminFilter),
			beego.NSInclude(
				&admin_controllers.AdminUsersController{},
				&admin_controllers.AdminAreasController{},
				&admin_controllers.AdminLocalsController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.ErrorController(&controllers.ErrorController{})
}
