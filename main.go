package main

import (
	"github.com/gorilla/context"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"fmt"
)

func main() {
	i18n := service.NewJsonI18n()
	err := i18n.ParseTranslationRoot("resource/translation")
	if err != nil {
		fmt.Printf("Could not install translations: %+v\n", err)
		return
	}
	home := handler.NewHome(i18n)
	http.HandleFunc("/", home.Index)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
