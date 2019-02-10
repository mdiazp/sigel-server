package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/area`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/area`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/area`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:AreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/local`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/local`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "Admins",
			Router: `/local/admins`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "PutAdmin",
			Router: `/local/admins`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "DeleteAdmin",
			Router: `/local/admins`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:LocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotification",
			Router: `/notification`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotifications",
			Router: `/notifications`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "Accept",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "Refuse",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/reservations`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "List2",
			Router: `/reservations2`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:ReservationsController"],
		beego.ControllerComments{
			Method: "ReservationsCount",
			Router: `/reservationscount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/user`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sigel-server/api/controllers/admin:UsersController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
