package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "markdown-blog/routers"
	"net/url"
	"os"
	"strings"
)

func main() {
	beego.AddFuncMap("add", Add)
	beego.AddFuncMap("addLeftPadZero", AddLeftPadZero)
	beego.AddFuncMap("stringArrayToString", StringArrayToString)
	beego.AddFuncMap("stringArrayToEscapeString", StringArrayToEscapeString)
	beego.SetStaticPath("/img", "static"+string(os.PathSeparator)+"img")
	beego.Run()
}
func Add(a int32, b int32) (result int32) {

	return a + b
}

func StringArrayToString(strArray []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(strArray), "[]"), " ", ",", -1)
}

func StringArrayToEscapeString(strArray []string) string {
	result := []string{}
	if len(strArray) > 0 {
		for _, str := range strArray {
			result = append(result, url.QueryEscape(str))
		}
	}
	if len(result) > 0 {
		return strings.Replace(strings.Trim(fmt.Sprint(result), "[]"), " ", ",", -1)
	} else {
		return ""
	}

}

func AddLeftPadZero(str int32) (result string) {
	s := string(str)
	s = "0000" + s
	return s[len(s)-4 : 4]
}
