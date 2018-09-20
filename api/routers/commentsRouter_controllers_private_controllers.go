package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers/private_controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/profile`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

}
