package common

import (
	"fmt"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetAbout() *Post {
	a := new(Post)

	filePathStr := "aboutPost" + string(os.PathSeparator) + "about.md"

	fileRead, fileReadErr := ioutil.ReadFile(filePathStr)

	if fileReadErr != nil {
		fmt.Printf(filePathStr)
		return nil
	}
	file := strings.Replace(filePathStr, "posts"+string(os.PathSeparator), "", -1)
	file = strings.Replace(file, ".md", "", -1)
	file = strings.Replace(file, ".markdown", "", -1)
	lines := strings.Split(string(fileRead), "\n")
	yamlPostInfo := ""
	bodyStartLine := 0
	for index, line := range lines {
		if strings.HasPrefix(line, "---") && index == 0 {
			continue
		}
		if yamlPostInfo == "" {
			yamlPostInfo = line
		} else {
			yamlPostInfo = yamlPostInfo + "\n" + line
		}
		if strings.HasPrefix(line, "---") && index > 0 {
			bodyStartLine = index + 1
			break
		}
	}
	if yamlPostInfo != "" {
		err := yaml.Unmarshal([]byte(yamlPostInfo), a)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
	body := strings.Join(lines[bodyStartLine:len(lines)], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	a.Body = body
	a.File = file
	//a.Next = GetPost()
	return a
}
