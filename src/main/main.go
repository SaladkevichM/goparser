package main

import (
	"html/template"
	"net/http"
	"parse"
	"time"
	"util"
)

var SOURCE = "http://tut.by/"
var TEMPLATE = "index.html"

var COOK_N = "goparser_n"
var COOK_T = "goparser_t"

type Page struct {
	BaseURL string
	Topics  *map[string]int
	News    *map[string]int
	Story   template.HTML
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/parser", http.StatusSeeOther)
}

func parserHandler(w http.ResponseWriter, r *http.Request) {
	url, story := SOURCE, ""

	t := util.AskCookie(COOK_T, r.FormValue("t"), r, w)
	n := util.AskCookie(COOK_N, r.FormValue("n"), r, w)

	topics := parse.Topics(util.Load("", url))
	news := parse.News(util.Load("", t))
	story = parse.Story(util.Load("", n))

	page := &Page{BaseURL: url, Topics: topics, News: news, Story: template.HTML(story)}
	renderTemplate(w, page)
}

func renderTemplate(w http.ResponseWriter, p *Page) {
	var t = template.Must(template.ParseFiles(TEMPLATE))
	err := t.ExecuteTemplate(w, TEMPLATE, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	_, err := client.Get("http://localhost:8080")
	if err != nil {
		http.HandleFunc("/", makeHandler(rootHandler))
		http.HandleFunc("/parser", makeHandler(parserHandler))
		http.ListenAndServe(":8080", nil)
	}

}
