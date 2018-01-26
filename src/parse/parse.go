package parse

import (
	"regexp"
	"strings"
	"util"
)

const CNP = "Couldn't parse"

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

// TODO. needs another way for extracting article
func Story(text string) string {
	parts := strings.Split(text, "class=\"js-mediator-article\">")
	if len(parts) > 1 {
		parts := strings.Split(parts[1], "</div>")
		if len(parts) > 1 {
			regexp, _ := regexp.Compile("<img.*?/>")
			return regexp.ReplaceAllString(parts[0], "")
		}
	}
	return CNP
}
