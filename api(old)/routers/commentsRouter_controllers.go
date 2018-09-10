package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:SessionController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:SessionController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
