package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/service"
)

func Router(i18n service.I18n, ad service.Ad) *mux.Router {
	r := mux.NewRouter()

	home := NewHome(i18n, ad)
	r.HandleFunc("/", home.Index)

	return r
}
