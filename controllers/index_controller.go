package controllers

import (
	"github.com/Kavin-Cao/rest-trains/models"
)

type IndexController struct {
	BaseController
}

// @router /index [get,post]
func (this *IndexController) Index() {
	sessionUser := this.GetSession(USER_SESSION_KEY)
	if sessionUser == nil {
		this.Redirect("/login", 302)
		return
	}
	var user models.User
	user,ok := sessionUser.(models.User)
	if !ok {
		this.Redirect("/login", 302)
		return
	}
	userId := user.ID
	//查询用户的项目
	var projects []models.Project
	models.RTDB.Where("Creator = ?", userId).Find(&projects)
	this.Data["userInfo"] = user
	this.Data["projects"] = projects
	this.TplName = "index.html"
}

// @router / [get,post]
func (this *IndexController) Index2() {
	this.Index()
}

// @router /login [get,post]
func (this *IndexController) Login() {
	if this.Ctx.Request.Method == "GET"{
		this.TplName = "login.html"
		return
	}
	var user models.User
	models.RTDB.First(&user, 1).Row()
	this.SetSession(USER_SESSION_KEY, user)
	this.Redirect("/index", 302)
}