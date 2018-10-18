package controllers

import (
	"github.com/astaxie/beego"
	"markdown-blog/common"
	"path/filepath"
	"sort"
)

type TagsController struct {
	beego.Controller
}

func (c *TagsController) Get() {
	//c.SetData()
	//:year/:mouth/:day/:postName
	//tag := c.Ctx.Input.Query("tag")
	//tag := c.Ctx.Input.Param(":tag")
	siteConfig := new(common.SiteConfig) // common.SiteConfig.GetConfig("config/_config.yaml")
	absPath, _ := filepath.Abs("./config/_config.yml")
	siteConfig.GetConfig(absPath)
	c.Data["site"] = siteConfig
	page := new(common.Post)

	result := common.GetTagsOrderInfo()
	tagOrders := []common.YearPosts{}
	//if tag == "" {
	page.HeaderImage = "static/img/tag-bg.jpg"
	page.Title = "Archive"
	c.Data["page"] = page
	c.TplName = "archive.html"
	c.Data["layout"] = "page"
	c.Layout = "layout/layout.html"
	for _, posts := range result {
		if posts != nil && len(posts) > 0 {
			for _, postArray := range posts {
				isContain := false
				for index, yearPosts := range tagOrders {
					if yearPosts.Year == int32(postArray.Year) {
						isContain = true
						for _, postTag := range postArray.Posts {
							isHavePost := false
							for _, orderPost := range yearPosts.Posts {
								if orderPost.Title == postTag.Title && orderPost.Date == postTag.Date {
									isHavePost = true
									break
								}
							}
							if !isHavePost {
								yearPosts.Posts = append(yearPosts.Posts, postTag)
								tagOrders[index] = yearPosts
							}
						}
					}
				}
				if !isContain {
					yearPosts := new(common.YearPosts)
					yearPosts.Year = int32(postArray.Year)
					yearPosts.Posts = postArray.Posts
					tagOrders = append(tagOrders, *yearPosts)
				}
			}
		}
	}
	tagOrderSlice := common.YearPostsSlice(tagOrders)
	sort.Stable(tagOrderSlice)
	sort.Sort(sort.Reverse(tagOrderSlice))
	c.Data["tagOrders"] = tagOrderSlice
	c.Data["tags"] = common.GetTags()
	//} else {
	//	if result[tag] == nil {
	//		post := new(common.Post)
	//		post.Layout = "page"
	//		post.Description = "你来到了没有知识的荒原 :("
	//		post.HeaderImage = "static/img/404-bg.jpg"
	//		c.Layout = "layout/layout.html"
	//		c.TplName = "404.html"
	//		c.Data["page"] = post
	//		c.Data["layout"] = post.Layout
	//	} else {
	//		c.Data["page"] = page
	//		page.Title = "Archive"
	//		page.HeaderImage = "static/img/tag-bg.jpg"
	//		c.TplName = "archive.html"
	//		c.Data["layout"] = "page"
	//		c.Data["tagOrders"] = result[tag]
	//		c.Data["tags"] = common.GetTags()
	//		c.Layout = "layout/layout.html"
	//	}

	//}
}
