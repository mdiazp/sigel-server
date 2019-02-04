package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:LogoutController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:LogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/session/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotification",
			Router: `/session/notification`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotifications",
			Router: `/session/notifications`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotificationsCount",
			Router: `/session/notificationscount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "SetUserNotificationsAsReaded",
			Router: `/session/readallnotifications`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "ReadNotification",
			Router: `/session/readnotification`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/session/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/session/profile`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/session/reservation`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"],
		beego.ControllerComments{
			Method: "Confirm",
			Router: `/session/reservation`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/session/reservations`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
