package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["myproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["myproject/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["myproject/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "AddArticle",
			Router:           "/",
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myproject/controllers:UserController"] = append(beego.GlobalControllerRouter["myproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
