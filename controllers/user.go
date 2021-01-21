package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {
	c.TplName = "user.tpl"
}

func (c *UserController) DoAddUser() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id必须是int类型")
		return
	}
	fmt.Printf("%v------%T", id, id)
	//这里和name和html中的name对应
	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetString("hobby")
	fmt.Printf("%v------%T", hobby, hobby)
	c.Ctx.WriteString("用户中心--" + strconv.Itoa(id) + username + password)
}
