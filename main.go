package main

import (
	"github.com/gorilla/context"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/util"
	"net/http"
	"fmt"
)

func main() {
	err := util.ParseTranslationDir("resource/translation")
	if err != nil {
		fmt.Printf("Could not install translations: %+v\n", err)
		return
	}
	handler.RouteInstall()
	util.TemplateInstall()
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
