package parse

import (
	"regexp"
	"strings"
	"util"
)

var phrases = []string{"utm_source", "www"}

func News(text string) *map[string]int {

	regexp, _ := regexp.Compile("(https://\\w+.tut.by/.*?/.*?html).*?")
	news := util.SliceToMap(regexp.FindAllString(text, -1), phrases)
	return news
}

func Topics(text string) *map[string]int {
	regexp, _ := regexp.Compile("(https://\\w+.tut.by/\\w+/).*?")
	topics := util.SliceToMap(regexp.FindAllString(text, -1), phrases)
	return topics
}

func Story(text string) string {
	regexp, _ := regexp.Compile(".*?article_body.*")
	return strings.Replace(regexp.FindString(text), "<strong>", "", -1) + "</div>"
}
