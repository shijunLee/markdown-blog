package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type SiteConfig struct {
	Title                 string   `yaml:"title"`
	SEOTitle              string   `yaml:"SEOTitle"`
	HeaderImage           string   `yaml:"header-img"`
	Email                 string   `yaml:"email"`
	Description           string   `yaml:"description"`
	Keyword               string   `yaml:"keyword"`
	Url                   string   `yaml:"url"`
	BaseUrl               string   `yaml:"baseurl"`
	Future                bool     `yaml:"future"`
	ChromeTabThemeColor   string   `yaml:"chrome-tab-theme-color"`
	SidebarAvatar         string   `yaml:"sidebar-avatar"`
	Friends               []Friend `yaml:"friends"`
	FeaturedTags          bool     `yaml:"featured-tags"`
	FeaturedConditionSize int32    `yaml:"featured-condition-size"`
	PostCount			  int32
    SidebarAboutDescription string `yaml:"sidebar-about-description"`
	ZhiHuUserName string `yaml:"zhihu_username"`
	WeiBoUserName string `yaml:"weibo_username"`
	GitHubUserName string `yaml:"github_username"`
}

type Friend struct {
	Title string `yaml:"title"`
	Href  string `yaml:"href"`
}

// 后期研究一下调用方法问题
func (c *SiteConfig) GetConfig(configPath string) *SiteConfig {

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	  postFiles,_ :=filepath.Glob("posts/*")
	  postCount :=len(postFiles)
	  c.PostCount = int32(postCount)
	return c
}
