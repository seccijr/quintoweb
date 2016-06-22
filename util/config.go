package util

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/service"
)

func Router(rootPath string, i18n service.I18n) *mux.Router {
	r := mux.NewRouter()

	home := handler.NewHome(rootPath, i18n)
	r.HandleFunc("/", home.Index)

	return r
}
