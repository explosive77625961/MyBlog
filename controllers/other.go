package controllers

import (
	"github.com/astaxie/beego"
)

type OtherControllers struct {
	beego.Controller
}

func (this *OtherControllers) Get(){
	this.TplName = "other.html"
	this.Data["IsOther"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

}
