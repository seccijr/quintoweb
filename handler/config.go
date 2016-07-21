package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"path"
	"golang.org/x/text/language"
)

type LangContextHandler struct {
	i18n service.I18n
}

func (h LangContextHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t, _, err := language.ParseAcceptLanguage(req.Header.Get("Accept-Language"))
	if err == nil {
		tag, _, _ := h.i18n.Match(t...)
		home := NewHome(i18n, ad, tag)
	}
	h.handler.ServeHTTP(w, req)
}

func NewLangContextHandler(i18n service.I18n) LangContextHandler {
	return LangContextHandler{i18n}
}

func Router(i18n service.I18n) *mux.Router {
	root := environment.Root()
	r := mux.NewRouter()

	publicHandler := http.StripPrefix("/public/", http.FileServer(http.Dir(path.Join(root, "public/"))))
	r.PathPrefix("/public/").Handler(publicHandler)

	localizedHome := NewLangContextHandler(i18n)
	r.Handle("/", localizedHome)

	return r
}
