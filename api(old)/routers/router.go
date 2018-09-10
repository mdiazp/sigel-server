// @APIVersion 1.0.0
// @Title api-sirel
// @Description Documentation for api-sirel
// @Contact di@upr.edu.cu
// @License UPR
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/filters"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "authHd"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	//beego.InsertFilter("/api/*", beego.BeforeRouter, filters.Auth)

	ns := beego.NewNamespace("/api/",
		beego.NSBefore(filters.Auth),

		beego.NSInclude(
			&controllers.SessionController{},
			// &controllers.ProfileController{},
			// &controllers.PublicAreasController{},
		),
		/*
			beego.NSNamespace("/admin",
				beego.NSBefore(filters.Admin),
				beego.NSInclude(
					&controllers.AdminUsersController{},
					&controllers.AdminAreasController{},
					&controllers.AdminLocalsController{},
				),
			),
		*/
	)
	beego.AddNamespace(ns)
}
