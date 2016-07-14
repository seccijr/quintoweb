package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"path"
	"golang.org/x/text/language"
	"fmt"
)

type LocalizedHandler interface {
	SetLanguage(lang language.Tag)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type LangContextHandler struct {
	handler LocalizedHandler
	i18n service.I18n
}

func (h LangContextHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t, _, err := language.ParseAcceptLanguage(req.Header.Get("Accept-Language"))
	if err == nil {
		tag, _, _ := h.i18n.Match(t...)
		h.handler.SetLanguage(tag)
	}
	h.handler.ServeHTTP(w, req)
}

func NewLangContextHandler(h LocalizedHandler, i18n service.I18n) LangContextHandler {
	return &LangContextHandler{h, i18n}
}

func Router(i18n service.I18n, ad service.Ad) *mux.Router {
	root := environment.Root()
	r := mux.NewRouter()

	publicHandler := http.StripPrefix("/public/", http.FileServer(http.Dir(path.Join(root, "public/"))))
	r.PathPrefix("/public/").Handler(publicHandler)

	home := NewHome(i18n, ad)
	localizedHome := NewLangContextHandler(home, i18n)
	r.HandleFunc("/", localizedHome)

	return r
}
