package handler

import (
	"net/http"
	"html/template"
	"github.com/seccijr/quintoweb/util"
)

func Index(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{
		"trans": util.GetTranslation,
	}
	t := template.New("index").Funcs(funcs)
	t, _ = t.ParseFiles("resource/view/layout.html")
	t.Execute(w, nil)
}
