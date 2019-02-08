package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:AreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:AreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:InfoController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:InfoController"],
		beego.ControllerComments{
			Method: "ServerTime",
			Router: `/servertime`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LoginController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:ReservationsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:ReservationsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/reservations`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:UsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:UsersController"],
		beego.ControllerComments{
			Method: "GetUserPublicInfo",
			Router: `/user/publicinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:UsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/public:UsersController"],
		beego.ControllerComments{
			Method: "GetUsersPublicInfo",
			Router: `/users/publicinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
