package routers

import (
	"github.com/astaxie/beego"
	"github.com/Kavin-Cao/rest-trains/controllers"
)

func init(){
	beego.Include(&controllers.IndexController{})
}