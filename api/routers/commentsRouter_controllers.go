package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"],
		beego.ControllerComments{
			Method: "CreateArea",
			Router: `/area`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"],
		beego.ControllerComments{
			Method: "GetArea",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"],
		beego.ControllerComments{
			Method: "UpdateArea",
			Router: `/area`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"] = append(beego.GlobalControllerRouter["github.com/mdiazp/sirel-server/api/controllers:TestingModel2Controller"],
		beego.ControllerComments{
			Method: "DeleteArea",
			Router: `/area`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
