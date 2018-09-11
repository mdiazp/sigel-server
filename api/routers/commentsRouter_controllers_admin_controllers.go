package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/area`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/area`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/area`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/local`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/local`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/user`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
