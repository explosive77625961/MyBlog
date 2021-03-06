package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Prepare(){
	this.Ctx.Output.Header("Cache-Control", "no-cache,no-store")
}

func (this *LoginController) Get() {
<<<<<<< HEAD
	if this.Input().Get("exit") == "true" {
		this.EnableRender = false
		maxAge := -1
		this.Ctx.SetCookie("uname", "", maxAge, "/")
		this.Ctx.SetCookie("pwd", "",maxAge, "/")

		this.Ctx.Redirect(301, "/")
		return
	}
=======

>>>>>>> 4c358306839b5e0b943ca9d17618fd030e98eda5

	if this.Input().Get("exit") == "true" {
		maxAge := -1
		this.Ctx.SetCookie("uname", "", maxAge, "/")
		this.Ctx.SetCookie("pwd", "",maxAge, "/")
		this.Ctx.Redirect(301, "/")
		return
	}
	this.TplName = "sigin.html"
}

func (this *LoginController) Post() {

	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	rember := this.Input().Get("rember") == "on"

	if uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd") {
		maxAge := 0
		if rember {
			maxAge = 1 << 31 -1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge,"/")
		this.Ctx.Redirect(301, "/")
		return
	}
	this.Ctx.Redirect(301, "/login")
	return
}


func checkAccount(c *context.Context) bool {
	ck, err := c.Request.Cookie("uname")
	if err != nil{
		return false
	}
	uname := ck.Value
	ck,err = c.Request.Cookie("pwd")
	if err != nil{
		return false
	}
	pwd := ck.Value
	return uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd")
}
