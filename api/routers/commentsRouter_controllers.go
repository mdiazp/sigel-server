package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/area`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/area/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/area/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/local/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/local/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:LoginController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:LogoutController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:ProfileController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:ProfileController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/profile`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/area/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
