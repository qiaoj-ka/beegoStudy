package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

// swagger注解配置
// @Title Get
// @Description get article 新闻列表
// @Success 200 {object} models
// @Failure 403 error
// @router / [Get]
func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表")
}

// swagger注解配置
// @Title Get
// @Description add article 增加新闻
// @Success 200 {object} models
// @Failure 403 error
// @router / [Get]
func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻")
}

// swagger注解配置
// @Title Get
// @Description add article 增加新闻
// @Success 200 {object} models
// @Failure 403 error
// @router / [Get]
func (c *ArticleController) EditArticle() {

	//获取 get 传值
	id := c.GetString("id")
	fmt.Printf("值:%v 类型:%T", id, id)
	c.Ctx.WriteString("修改新闻" + id)
}
