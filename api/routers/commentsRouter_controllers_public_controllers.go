package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicLocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicLocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicReservationsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicUsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/public_controllers:PublicUsersController"],
		beego.ControllerComments{
			Method: "GetUsernamesList",
			Router: `/users/usernames`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
