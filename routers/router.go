// @APIVersion 1.0.1
// @Title beego Test API for 自定义
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact jiajia@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// 配置路由映射,以上注解必须添加
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// ArticleController的路由映射
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		// 其余Controller的路由映射在这里添加...
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/articleAdd",
			beego.NSRouter("/",
				&controllers.ArticleController{}, "get:AddArticle"),
		),
		beego.NSNamespace("/articleEdit",
			beego.NSRouter("/",
				&controllers.ArticleController{}, "get:EditArticle"),
		),
		beego.NSNamespace("user/doAddUser",
			beego.NSRouter("/",
				&controllers.UserController{}, "post:DoAddUser"),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
	//Get路由,获取新闻
	//beego.Router("/article", &controllers.ArticleController{})
	//Get路由，增加新闻,访问自定义
	//beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	//Get路由，修改新闻，结构体地址
	//beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")
	//Get路由，获取用户
	//beego.Router("/user", &controllers.UserController{})
	//Get路由，添加用户
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	//Post路由，添加用户
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
}
