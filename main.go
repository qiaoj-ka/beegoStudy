package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "myproject/routers"
)

func main() {
	//=======在Terminal内输入bee run -gendoc=true -downdoc=true命令=====
	//然后在配置如下即可启动swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//======以上是swagger===============================================
	beego.Run()
}
