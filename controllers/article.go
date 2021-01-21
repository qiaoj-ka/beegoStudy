package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻")
}

func (c *ArticleController) EditArticle() {

	//获取 get 传值
	id := c.GetString("id")
	fmt.Printf("值:%v 类型:%T", id, id)
	c.Ctx.WriteString("修改新闻" + id)
}
