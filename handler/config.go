package handler

import (
	"github.com/gorilla/mux"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"github.com/seccijr/quintoweb/environment"
	"path"
)

func Router(i18n service.I18n, ad service.Ad) *mux.Router {
	root := environment.Root()
	r := mux.NewRouter()

	publicHandler := http.StripPrefix("/public/", http.FileServer(http.Dir(path.Join(root, "public/"))))
	r.PathPrefix("/public/").Handler(publicHandler)

	home := NewHome(i18n, ad)
	r.HandleFunc("/", home.Index)

	return r
}
