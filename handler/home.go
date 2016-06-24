package handler

import (
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/model"
	"github.com/seccijr/quintoweb/service"
	"golang.org/x/text/language"
	"html/template"
	"net/http"
	"path/filepath"
)

type Home struct {
	i18n service.I18n
	ad   service.Ad
	lang language.Tag
}

func NewHome(i18n service.I18n, ad service.Ad) Home {
	return Home{i18n, ad}
}

func (home Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	root := environment.Root()
	templateFiles := []string{
		filepath.Join(root, "resource/view/layout/base.html"),
		filepath.Join(root, "resource/view/home/index.html"),
	}
	tName := filepath.Base(templateFiles[0])
	funcs := template.FuncMap{
		"trans": home.i18n.TransTemplate,
	}
	b := model.Base{
		"indexTitle",
		[]byte(""),
		home.lang,
	}
	p := &model.Index{
		b,
		home.ad.GetTopDescOrder(10),
	}
	t, _ := template.New(tName).Funcs(funcs).ParseFiles(templateFiles...)
	t.ExecuteTemplate(w, tName, p)
}

func (home Home) SetLanguage(lang language.Tag) {
	home.lang = lang
}
