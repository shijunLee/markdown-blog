package controllers

import (
	"github.com/astaxie/beego"
	"markdown-blog/common"
	"path/filepath"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Get() {
	//engine := liquid.NewEngine()
	//template := `<h1>{{ page.title }}</h1>`
	//bindings := map[string]interface{}{
	//	"page": map[string]string{
	//		"title": "Introduction",
	//	},
	//}
	//out, err := engine.ParseAndRenderString(template, bindings)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//layout: default
	//title: 404
	//hide-in-nav: true
	//description: "你来到了没有知识的荒原 :("
	//header-img: "/static/img/404-bg.jpg"
	//permalink: /404.html
	post := new(common.Post)
	post.Layout = "page"
	post.Description = "你来到了没有知识的荒原 :("
	post.HeaderImage = "static/img/404-bg.jpg"
	this.Layout = "layout/layout.html"
	this.TplName = "404.html"
	this.Data["page"] = post
	this.Data["layout"] = post.Layout
	siteConfig := new(common.SiteConfig) // common.SiteConfig.GetConfig("config/_config.yaml")
	absPath, _ := filepath.Abs("./config/_config.yml")
	siteConfig.GetConfig(absPath)
	this.Data["site"] = siteConfig
	//this.LayoutSections["HtmlHeader"]="_includes/intro-header.html"
}
