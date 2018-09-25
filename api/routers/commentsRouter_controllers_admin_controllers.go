package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/area`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/area`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/area`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/local`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/local`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminLocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/user`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/admin_controllers:AdminUsersController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
