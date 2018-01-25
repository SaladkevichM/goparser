package util

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SliceToMap(slice []string, skip []string) *map[string]int {
	m := make(map[string]int)
	for _, v := range slice {
		m[v] = 1
	}
	for k, _ := range m {
		for _, v := range skip {
			if strings.Contains(k, v) {
				delete(m, k)
			}
		}
	}
	return &m
}

func AskCookie(name string, value string, r *http.Request, w http.ResponseWriter) string {

	expire := time.Now().Add(365 * (time.Hour * 24))
	newCookie := http.Cookie{Name: name, Value: value, Path: "/", Expires: expire}
	cookie, noCookie := r.Cookie(name)

	if noCookie != nil {
		http.SetCookie(w, &newCookie)
		return value
	} else {
		if cookie.Value != value && value != "" {
			http.SetCookie(w, &newCookie)
			return value
		}
		return cookie.Value
	}

}

func Load(t string, url string) string {

	r := ""
	if url == "" {
		return r
	}

	switch t {
	case "topics":
		f, _ := ioutil.ReadFile("test_sources/topics.html")
		r = string(f)
	case "news":
		f, _ := ioutil.ReadFile("test_sources/news.html")
		r = string(f)
	case "story":
		f, _ := ioutil.ReadFile("test_sources/story.html")
		r = string(f)
	default:
		r = getString(url)
	}

	return r
}

func getString(url string) string {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}
