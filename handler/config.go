package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"path"
)

type LangContextHandler struct {
	i18n service.I18n `inject:""`
	handler http.Handler
}

func (h LangContextHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//t, _, err := language.ParseAcceptLanguage(req.Header.Get("Accept-Language"))
	//if err == nil {
	//	tag, _, _ := h.i18n.Match(t...)
	//}
	h.handler.ServeHTTP(w, req)
}

func NewLangContextHandler(handler http.Handler) LangContextHandler {
	var ctxHandler LangContextHandler
	ctxHandler.handler = handler

	return ctxHandler
}

func Router() *mux.Router {
	root := environment.Root()
	r := mux.NewRouter()

	publicHandler := http.StripPrefix("/public/", http.FileServer(http.Dir(path.Join(root, "public/"))))
	r.PathPrefix("/public/").Handler(publicHandler)

	localizedHome := NewLangContextHandler(Home{})
	r.Handle("/", localizedHome)

	return r
}
