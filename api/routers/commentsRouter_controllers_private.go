package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:LogoutController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:LogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotification",
			Router: `/profile/notification`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:NotificationsController"],
		beego.ControllerComments{
			Method: "GetNotifications",
			Router: `/profile/notification`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ProfileController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/profile`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private:ReservationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/reservation`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
