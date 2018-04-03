package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Index2",
			Router: `/`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/index`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/Kavin-Cao/rest-trains/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

}
