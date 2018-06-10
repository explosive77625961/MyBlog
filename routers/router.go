package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/other", &controllers.OtherControllers{})
	beego.Router("/login", &controllers.LoginController{})
}
