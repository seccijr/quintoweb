package handler

import (
	"net/http"
	"html/template"
	"golang.org/x/text/language"
	"github.com/seccijr/quintoweb/model"
	"path/filepath"
	"github.com/seccijr/quintoweb/service"
)

type Home struct {
	i18n service.I18n
}

func NewHome(i18n service.I18n) Home {
	return Home{i18n}
}

func (home Home) Index(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"resource/view/layout/base.html",
		"resource/view/home/index.html",
	}
	tName := filepath.Base(templateFiles[0])
	funcs := template.FuncMap{
		"trans": home.i18n.TransTemplate,
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
