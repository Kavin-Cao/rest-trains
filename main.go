package main

import (
	_ "github.com/Kavin-Cao/rest-trains/routers"
	_ "github.com/Kavin-Cao/rest-trains/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

