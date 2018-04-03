package controllers

import "github.com/astaxie/beego"

const USER_SESSION_KEY  = "$USER:"

type BaseController struct {
	beego.Controller
}