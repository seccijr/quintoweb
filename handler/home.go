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
}

func NewHome(i18n service.I18n) Home {
	return Home{i18n}
}

func (home Home) Index(w http.ResponseWriter, r *http.Request) {
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
		language.MustParse("es"),
	}
	p := &model.Index{
		b,
		service.GetTopAdsDescOrder(10),
	}
	t, _ := template.New(tName).Funcs(funcs).ParseFiles(templateFiles...)
	t.ExecuteTemplate(w, tName, p)
}
