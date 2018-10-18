package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"markdown-blog/common"
	"strings"
)

type PostDetailController struct {
	beego.Controller
}

func (c *PostDetailController) Get() {
	//c.SetData()
	//:year/:mouth/:day/:postName
	year := c.Ctx.Input.Param(":year")
	mouth := c.Ctx.Input.Param(":mouth")
	day := c.Ctx.Input.Param(":day")
	postName := c.Ctx.Input.Param(":postName")
	fileName := postName
	if year != "" && mouth != "" && day != "" {
		fileName = fmt.Sprintf("%v-%v-%v-%v", year, mouth, day, postName)
	}
	post := common.GetPost(fileName, false)

	c.TplName = "post.html"
	c.Layout = "layout/layout-post.html"
	c.Data["layout"] = "post"
	if strings.ToLower(post.Layout) == "post" {
		c.Layout = "layout/layout-post.html"
	} else if strings.ToLower(post.Layout) == "keynote" {
		c.Layout = "layout/layout-keynote.html"
		c.Data["layout"] = "keynote"
	}
	c.Data["tags"] = common.GetTags()
	c.Data["page"] = post
}
