package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"markdown-blog/common"
	"path/filepath"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.SetData()
	pageIndex := c.Ctx.Input.Param(":pageIndex")

	if pageIndex == "" {
		pageIndex = "page1"
	}

	c.Data["posts"] = common.GetPosts()
	c.TplName = "page.html"
	siteConfig := new(common.SiteConfig) // common.SiteConfig.GetConfig("config/_config.yaml")
	absPath, _ := filepath.Abs("./config/_config.yml")
	siteConfig.GetConfig(absPath)
	c.Data["site"] = siteConfig
	//post := new(common.Post)
	//post.Layout = "page"
	//post.HeaderImage=""
	posts := common.GetPosts()
	indexStr := pageIndex[4:]
	currentPage, error := strconv.Atoi(indexStr)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	if currentPage <= 1 {
		c.Data["prePage"] = ""
	} else {
		c.Data["prePage"] = fmt.Sprintf("page%d", currentPage-1)
	}
	if (currentPage-1)*10+10 < len(posts) {
		posts = posts[(currentPage-1)*10 : currentPage*10]
		c.Data["nextPage"] = fmt.Sprintf("page%d", currentPage+1)
	} else {
		posts = posts[(currentPage-1)*10:]
		c.Data["nextPage"] = ""
	}
	c.Layout = "layout/layout-post.html"
	page := new(common.Post)
	c.Data["tags"] = common.GetTags()
	c.Data["page"] = page
	c.Data["posts"] = posts
	c.Data["layout"] = "page"
}
