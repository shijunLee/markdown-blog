package common

import (
	"fmt"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Post struct {
	Layout                string   `yaml:"layout"`
	Title                 string   `yaml:"title"`
	SubTitle              string   `yaml:"subtitle"`
	Author                string   `yaml:"author"`
	Tags                  []string `yaml:"tags"`
	HeaderImage           string   `yaml:"header-img"`
	HeaderMask            float64  `yaml:"header-mask"`
	HeaderBGCss           string   `yaml:"header-bg-css"`
	HeaderStyle           string   `yaml:"header-style"`
	HeaderImageCredit     string   `yaml:"header-img-credit"`
	HeaderImageCreditHref string   `yaml:"header-img-credit-href"`
	IFrame                string   `yaml:"iframe"`
	Description           string   `yaml:"description"`
	Short                 bool
	Date                  string
	Summary               string
	Body                  string
	File                  string
	Url                   string
	Catalog               string
	Multilingual          bool
	NavStyle              string `yaml:"nav-style"`
	Previous              *Post
	Next                  *Post
	DateTime              time.Time
}

type PostSlice []Post

func (s PostSlice) Len() int           { return len(s) }
func (s PostSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PostSlice) Less(i, j int) bool { return s[i].DateTime.Before(s[j].DateTime) }

type YearPostsSlice []YearPosts

func (s YearPostsSlice) Len() int           { return len(s) }
func (s YearPostsSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s YearPostsSlice) Less(i, j int) bool { return s[i].Year < s[j].Year }

type TagInfo struct {
	TagName string
	Count   int32
}

type YearPosts struct {
	Year  int32
	Posts PostSlice
}

//map[string]map[int32]PostSlice
type TagOrderInfo struct {
	tag       string
	YearPosts YearPostsSlice
}

func GetTagsOrderInfo() map[string]YearPostsSlice {
	tagOrderInfos := make(map[string]YearPostsSlice)
	posts := GetPosts()
	for _, post := range posts {
		tags := post.Tags
		if len(tags) > 0 {
			for _, tag := range tags {
				if tagOrderInfos[tag] == nil {

					yearPost := new(YearPosts)
					yearPost.Year = int32(post.DateTime.Year())
					tagPosts := []Post{}
					tagPostsSlice := PostSlice(tagPosts)
					tagPostsSlice = append(tagPostsSlice, post)
					sort.Stable(tagPostsSlice)
					sort.Sort(sort.Reverse(tagPostsSlice))
					yearPost.Posts = tagPostsSlice
					//yearPostsSlice := YearPostsSlice(yearPosts)
					tagOrderInfos[tag] = []YearPosts{}
					tagOrderInfos[tag] = append(tagOrderInfos[tag], *yearPost)
				} else {
					isContainYear := false
					yearPostsIndex := -1
					for index, yearPosts := range tagOrderInfos[tag] {
						if yearPosts.Year == int32(post.DateTime.Year()) {
							isContainYear = true
							yearPostsIndex = index
							break
						}
					}
					if isContainYear {
						tagOrderInfos[tag][yearPostsIndex].Posts = append(tagOrderInfos[tag][yearPostsIndex].Posts, post)
						sort.Stable(tagOrderInfos[tag][yearPostsIndex].Posts)
						sort.Sort(sort.Reverse(tagOrderInfos[tag][yearPostsIndex].Posts))
					} else {
						yearPost := new(YearPosts)
						yearPost.Year = int32(post.DateTime.Year())
						tagPosts := []Post{}
						tagPostsSlice := PostSlice(tagPosts)
						tagPostsSlice = append(tagPosts, post)
						sort.Stable(tagPostsSlice)
						sort.Sort(sort.Reverse(tagPostsSlice))
						yearPost.Posts = tagPostsSlice
						tagOrderInfos[tag] = append(tagOrderInfos[tag], *yearPost)
						sort.Stable(tagOrderInfos[tag])
						sort.Sort(sort.Reverse(tagOrderInfos[tag]))
					}
				}
			}
		}
	}
	return tagOrderInfos
}

func GetTags() []TagInfo {
	tags := []TagInfo{}
	files, _ := filepath.Glob("posts/*")
	for _, f := range files {
		post := new(Post)
		file := strings.Replace(f, "posts"+string(os.PathSeparator), "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileRead, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileRead), "\n")
		yamlPostInfo := ""
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
				break
			}
		}
		if yamlPostInfo != "" {
			err := yaml.Unmarshal([]byte(yamlPostInfo), post)
			if err != nil {
				log.Fatalf("Unmarshal: %v", err)
			}
		}
		if len(post.Tags) > 0 {
			for _, tag := range post.Tags {
				fmt.Println(tag)
				tagsContain := false
				for _, tagInfo := range tags {
					if tagInfo.TagName == tag {
						tagsContain = true
						tagInfo.Count = tagInfo.Count + 1
					}
				}
				if !tagsContain {
					tagInfo := new(TagInfo)
					tagInfo.TagName = tag
					tagInfo.Count = 1
					tags = append(tags, *tagInfo)
				}
			}
		}
	}
	return tags
}

func getPostUrl(fileName string) string {

	fileExtIndex := strings.LastIndex(fileName, ".")
	if fileExtIndex > 0 {
		fileName = fileName[:fileExtIndex]
	}
	if fileName == "" {
		return fileName
	}
	if len(fileName) < 11 {
		return fileName
	}
	dateStr := fileName[:10]
	postName := fileName[11:]
	return fmt.Sprintf("%v/%v", strings.Replace(dateStr, "-", "/", -1), postName)
}

func GetPosts() PostSlice {
	a := []Post{}
	files, _ := filepath.Glob("posts/*")
	for _, f := range files {
		post := new(Post)
		file := strings.Replace(f, "posts"+string(os.PathSeparator), "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Println(err)
		}
		createDateTime := fileInfo.ModTime()
		createDate := fmt.Sprintf("%v %d,%d", fileInfo.ModTime().Month().String(), fileInfo.ModTime().Day(), fileInfo.ModTime().Year())
		fileRead, _ := ioutil.ReadFile(f)
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
			err = yaml.Unmarshal([]byte(yamlPostInfo), post)
			if err != nil {
				log.Fatalf("Unmarshal: %v", err)
			}
		}
		markdownBody := strings.Join(lines[bodyStartLine:len(lines)], "\n")
		body := string(blackfriday.MarkdownCommon([]byte(markdownBody)))
		post.File = file
		post.Url = getPostUrl(file)
		post.Body = body
		fileNameDate, fileNameDateTime := getPostDate(fileInfo.Name())
		if fileNameDate != "" {
			post.DateTime = fileNameDateTime
			post.Date = fileNameDate
		} else {
			post.Date = createDate
			post.DateTime = createDateTime
		}
		rs := []rune(markdownBody)
		length := len(rs)
		if length == 0 {
			post.Summary = ""
		} else if length < 400 {
			post.Summary = string(rs[0:length])
		} else {
			post.Summary = string(rs[0:400]) + " ..."
		}
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		post.Summary = re.ReplaceAllString(post.Summary, " ")
		a = append(a, *post)
	}
	postSlices := PostSlice(a)
	sort.Stable(postSlices)
	sort.Sort(sort.Reverse(postSlices))

	return postSlices
}

func getFilePath(postName string) string {
	pattern := "posts" + string(os.PathSeparator) + postName + ".md"
	//dir, _ := filepath.Split(pattern)
	//if runtime.GOOS == "windows" {
	//	_, dir = cleanGlobPathWindows(dir)
	//} else {
	//	dir = cleanGlobPath(dir)
	//}

	_, err := os.Stat(pattern)
	if err != nil {
		pattern = "posts" + string(os.PathSeparator) + postName + ".markdown"
		_, err = os.Stat(pattern)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return pattern
}

func GetPost(postName string, isGetSub bool) *Post {
	a := new(Post)
	files, _ := filepath.Glob("posts/*")
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	filePathStr := getFilePath(postName)

	fileRead, fileReadErr := ioutil.ReadFile(filePathStr)

	if fileReadErr != nil {
		fmt.Printf(filePathStr)
		return nil
	}

	fileInfo, err := os.Stat(filePathStr)
	if err != nil {
		fmt.Println(err)
	}
	createDate := fmt.Sprintf("%v %d,%d", fileInfo.ModTime().Month().String(), fileInfo.ModTime().Day(), fileInfo.ModTime().Year())
	file := strings.Replace(filePathStr, "posts"+string(os.PathSeparator), "", -1)
	file = strings.Replace(file, ".md", "", -1)
	file = strings.Replace(file, ".markdown", "", -1)
	currentIndex := 0
	for index, fileName := range files {
		if strings.Contains(fileName, file) {
			currentIndex = index
		}
	}
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
		err = yaml.Unmarshal([]byte(yamlPostInfo), a)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
	body := strings.Join(lines[bodyStartLine:len(lines)], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	a.Body = body
	a.File = file
	a.Url = getPostUrl(file)
	fileNameDate, _ := getPostDate(fileInfo.Name())
	if fileNameDate != "" {
		a.Date = fileNameDate
	} else {
		a.Date = createDate
	}
	if !isGetSub {
		if currentIndex <= len(files)-2 {
			nextFileName := files[currentIndex+1]
			file := strings.Replace(nextFileName, "posts"+string(os.PathSeparator), "", -1)
			file = strings.Replace(file, ".md", "", -1)
			file = strings.Replace(file, ".markdown", "", -1)
			a.Next = GetPost(file, true)
		}

		if currentIndex > 0 {
			preFileName := files[currentIndex-1]
			file := strings.Replace(preFileName, "posts"+string(os.PathSeparator), "", -1)
			file = strings.Replace(file, ".md", "", -1)
			file = strings.Replace(file, ".markdown", "", -1)
			a.Previous = GetPost(file, true)
		}
	}
	//a.Next = GetPost()
	return a
}

func getPostDate(postName string) (date string, dateTime time.Time) {
	if postName == "" {
		return "", time.Time{}
	}
	if len(postName) < 11 {
		return "", time.Time{}
	}
	dateStr := postName[:10] + " 00:00:00"
	parseStrTime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		fmt.Println(err)
		return "", time.Time{}
	}
	dateStr = fmt.Sprintf("%v %d,%d", parseStrTime.Month().String(), parseStrTime.Day(), parseStrTime.Year()) //parseStrTime.Format("September 27, 2018")
	return dateStr, parseStrTime
}

// cleanGlobPath prepares path for glob matching.
func cleanGlobPath(path string) string {
	switch path {
	case "":
		return "."
	case string(os.PathSeparator):
		// do nothing to the path
		return path
	default:
		return path[0 : len(path)-1] // chop off trailing separator
	}
}

func cleanGlobPathWindows(path string) (prefixLen int, cleaned string) {
	vollen := volumeNameLen(path)
	switch {
	case path == "":
		return 0, "."
	case vollen+1 == len(path) && os.IsPathSeparator(path[len(path)-1]): // /, \, C:\ and C:/
		// do nothing to the path
		return vollen + 1, path
	case vollen == len(path) && len(path) == 2: // C:
		return vollen, path + "." // convert C: into C:.
	default:
		if vollen >= len(path) {
			vollen = len(path) - 1
		}
		return vollen, path[0 : len(path)-1] // chop off trailing separator
	}
}

// volumeNameLen returns length of the leading volume name on Windows.
// It returns 0 elsewhere.
func volumeNameLen(path string) int {
	if len(path) < 2 {
		return 0
	}
	// with drive letter
	c := path[0]
	if path[1] == ':' && ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
		return 2
	}
	// is it UNC? https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx
	if l := len(path); l >= 5 && isSlash(path[0]) && isSlash(path[1]) &&
		!isSlash(path[2]) && path[2] != '.' {
		// first, leading `\\` and next shouldn't be `\`. its server name.
		for n := 3; n < l-1; n++ {
			// second, next '\' shouldn't be repeated.
			if isSlash(path[n]) {
				n++
				// third, following something characters. its share name.
				if !isSlash(path[n]) {
					if path[n] == '.' {
						break
					}
					for ; n < l; n++ {
						if isSlash(path[n]) {
							break
						}
					}
					return n
				}
				break
			}
		}
	}
	return 0
}

func isSlash(c uint8) bool {
	return c == '\\' || c == '/'
}
