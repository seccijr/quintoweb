package handler

import (
	"net/http"
	"html/template"
	"golang.org/x/text/language"
	"github.com/seccijr/quintoweb/model"
	"path/filepath"
	"fmt"
	"github.com/seccijr/quintoweb/util"
)

func Test(args... interface{}) string {
	fmt.Printf("Called")
	return "Test"
}

func Index(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"resource/view/layout/base.html",
		"resource/view/home/index.html",
	}
	tName := filepath.Base(templateFiles[0])
	funcs := template.FuncMap{
		"trans": util.TransTemplate,
	}
	p := &model.Page{
		"indexTitle",
		[]byte(""),
		language.MustParse("es"),
	}
	t, _ := template.New(tName).Funcs(funcs).ParseFiles(templateFiles...)
	t.ExecuteTemplate(w, tName, p)
}
