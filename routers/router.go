package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//基本Get路由
}
