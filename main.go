package main

import (
	"fmt"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"path/filepath"
)

func main() {
	root := environment.Root()
	i18n := service.NewJsonI18n()
	err := i18n.ParseTranslationRoot(filepath.Join(root, "resource/translation"))
	if err != nil {
		fmt.Printf("Could not install translations: %+v\n", err)
		return
	}
	r := handler.Router(i18n)
	http.ListenAndServe(":8080", r)
}
