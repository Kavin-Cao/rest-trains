package routers

import (
	"github.com/Kavin-Cao/rest-trains/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
}
