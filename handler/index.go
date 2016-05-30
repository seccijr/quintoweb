package handler

import (
	"net/http"
	"html/template"
	"golang.org/x/text/language"
	"github.com/seccijr/quintoweb/model"
	"path/filepath"
	"github.com/seccijr/quintoweb/util"
	"github.com/seccijr/quintoweb/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"resource/view/layout/base.html",
		"resource/view/home/index.html",
	}
	tName := filepath.Base(templateFiles[0])
	funcs := template.FuncMap{
		"trans": util.TransTemplate,
	}
	b := model.Base{
		"indexTitle",
		[]byte(""),
		language.MustParse("es"),
	}
	p := &model.Index{
		b,
		service.GetTopAdsDescOrder(10),
	}
	t, _ := template.New(tName).Funcs(funcs).ParseFiles(templateFiles...)
	t.ExecuteTemplate(w, tName, p)
}
