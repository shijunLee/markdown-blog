package controllers

import (
	"github.com/astaxie/beego"
	"markdown-blog/common"
	"path/filepath"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {

	post, siteConfig := getErrorContent()
	c.TplName = "404.html"
	// c.Layout = "layout/layout-post.html"
	c.Data["page"] = post
	c.Data["layout"] = "page"
	c.Data["site"] = siteConfig
}

func (c *ErrorController) Error500() {

	post, siteConfig := getErrorContent()
	c.TplName = "404.html"
	// c.Layout = "layout/layout-post.html"
	c.Data["page"] = post
	c.Data["layout"] = "page"
	c.Data["site"] = siteConfig
}

func (c *ErrorController) Error501() {

	post, siteConfig := getErrorContent()
	c.TplName = "404.html"
	// c.Layout = "layout/layout-post.html"
	c.Data["page"] = post
	c.Data["layout"] = "page"
	c.Data["site"] = siteConfig
}

func (c *ErrorController) Error400() {

	post, siteConfig := getErrorContent()
	c.TplName = "404.html"
	// c.Layout = "layout/layout-post.html"
	c.Data["page"] = post
	c.Data["layout"] = "page"
	c.Data["site"] = siteConfig
}

func (c *ErrorController) Error401() {

	post, siteConfig := getErrorContent()
	c.TplName = "404.html"
	// c.Layout = "layout/layout-post.html"
	c.Data["page"] = post
	c.Data["layout"] = "page"
	c.Data["site"] = siteConfig
}

func getErrorContent() (postResult *common.Post, siteConfigResult *common.SiteConfig) {
	post := new(common.Post)
	post.Layout = "page"
	post.Description = "你来到了没有知识的荒原 :("
	post.HeaderImage = "static/img/404-bg.jpg"
	siteConfig := new(common.SiteConfig) // common.SiteConfig.GetConfig("config/_config.yaml")
	absPath, _ := filepath.Abs("./config/_config.yml")
	siteConfig.GetConfig(absPath)
	return post, siteConfig
}
