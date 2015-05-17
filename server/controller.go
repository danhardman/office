package server

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/danhardman/officr/server/pages"
)

// GetIndex displays the index page
func GetIndex(w http.ResponseWriter, r *http.Request) {
	p := &pages.Page{
		Title: "Officr",
	}

	t, err := template.ParseFiles("../static/index.html", "../static/header.html", "../static/footer.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, p)
}

func UpdateTemp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	b := r.Form
	t := b["temperature"][0]

	f, err := filepath.Abs("../temperature")

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(f, []byte(t), 0644)
	http.Redirect(w, r, "/", 301)
}
