package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/service"
)

func Router(i18n service.I18n) *mux.Router {
	r := mux.NewRouter()

	home := NewHome(i18n)
	r.HandleFunc("/", home.Index)

	return r
}
