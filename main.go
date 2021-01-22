package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "myproject/routers"
	"strings"
)

//这里应该加载 ""github.com/beego/beego/v2/server/web/context"" 否则会加载src/context
//过滤器实现
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("id").(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
	if !ok && !ok2 {
		ctx.Redirect(302, "/login/index")
	}
}

func main() {
	//=======在Terminal内输入bee run -gendoc=true -downdoc=true命令=====
	//然后在配置如下即可启动swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//======以上是swagger===============================================
	//======以下是验证用户是否登录过滤器=================================
	//注册过滤器
	beego.InsertFilter("*/", beego.BeforeRouter, FilterUser)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//======学习网页https://studygolang.com/articles/16395?fr=sidebar====
	beego.Run()
}
