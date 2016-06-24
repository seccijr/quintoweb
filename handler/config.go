package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"path"
	"golang.org/x/text/language"
)

type LocalizedHandler interface {
	SetLanguage(lang language.Tag)
}

type LangContextHandler struct {
	handler LocalizedHandler
}

func (h LangContextHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	headers := req.Header
}

func Router(i18n service.I18n, ad service.Ad) *mux.Router {
	root := environment.Root()
	r := mux.NewRouter()

	publicHandler := http.StripPrefix("/public/", http.FileServer(http.Dir(path.Join(root, "public/"))))
	r.PathPrefix("/public/").Handler(publicHandler)

	home := NewHome(i18n, ad)
	r.HandleFunc("/", home)

	return r
}
