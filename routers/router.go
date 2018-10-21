package routers

import (
	"github.com/astaxie/beego"
	"markdown-blog/controllers"
)

func init() {
	beego.Router("/:pageIndex", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post/:year/:mouth/:day/:postName", &controllers.PostDetailController{})
	beego.Router("/post/:postName", &controllers.PostDetailController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/archive", &controllers.TagsController{})
	beego.ErrorController(&controllers.ErrorController{})
}
