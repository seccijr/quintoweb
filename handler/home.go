package handler

import (
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/model"
	"github.com/seccijr/quintoweb/service"
	"html/template"
	"net/http"
	"path/filepath"
)

type Home struct {
	i18n service.I18n  `inject:""`
	ad   service.Ad  `inject:""`
}

func (home Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	root := environment.Root()
	templateFiles := []string{
		filepath.Join(root, "resource/view/layout/base.html"),
		filepath.Join(root, "resource/view/home/index.html"),
	}
	tName := filepath.Base(templateFiles[0])
	funcs := template.FuncMap{
		"trans": home.i18n.TaggedTransTemplateFunc(),
	}
	p := &model.Index{
		model.Base{
			Title: "indexTitle",
			Body: []byte(""),
			Lang: home.lang,
		},
		Ads: home.ad.GetTopDescOrder(10),
	}
	t, _ := template.New(tName).Funcs(funcs).ParseFiles(templateFiles...)
	t.ExecuteTemplate(w, tName, p)
}
