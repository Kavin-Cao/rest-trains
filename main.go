package main

import (
	_ "github.com/Kavin-Cao/rest-trains/models"
	_ "github.com/Kavin-Cao/rest-trains/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)
var globalSessions *session.Manager
func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}
func main() {
	beego.Run()
}

