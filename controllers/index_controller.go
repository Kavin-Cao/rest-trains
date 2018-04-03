package controllers

type IndexController struct {
	BaseController
}

// @Router / [get,post]
func (this *IndexController) Get() {
	user := this.GetSession(USER_SESSION_KEY)
	if user == nil {
		// write flash message
		//this.FlashWrite("NotPermit", "true")
		this.Redirect("/login", 302)
		return
	}
}
