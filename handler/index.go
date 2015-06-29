package handler

import (
	"github.com/seccijr/quintoweb/util"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	p := &util.Page{}
	renderTemplate(w, "index", p)
}

func View(w http.ResponseWriter, r *http.Request, title string) {
	p, err := util.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func Edit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := util.LoadPage(title)
	if err != nil {
		p = &util.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func Save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &util.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *util.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
