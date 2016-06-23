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

	cssh := http.StripPrefix("/css/", http.FileServer(http.Dir(path.Join(root, "public/css/"))))
	r.PathPrefix("/css/").Handler(cssh)

	vendorh := http.StripPrefix("/vendor/", http.FileServer(http.Dir(path.Join(root, "public/vendor/"))))
	r.PathPrefix("/vendor/").Handler(vendorh)

	home := NewHome(i18n, ad)
	r.HandleFunc("/", home.Index)

	return r
}
