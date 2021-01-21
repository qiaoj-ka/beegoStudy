package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//Get路由,获取新闻
	beego.Router("/article", &controllers.ArticleController{})
	//Get路由，增加新闻,访问自定义
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	//Get路由，修改新闻，结构体地址
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")
	//Get路由，获取用户
	beego.Router("/user", &controllers.UserController{})
	//Get路由，添加用户
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	//Post路由，添加用户
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
}
