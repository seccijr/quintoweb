package main

import (
	"fmt"
	"net/http"
	"github.com/seccijr/quintoweb/service"
	"github.com/seccijr/quintoweb/util"
	"os"
)

func main() {
	rootPath := os.Getenv("QUINTO_PATH");
	if rootPath  == "" {
		rootPath  = "/etc/root"
	}
	i18n := service.NewJsonI18n()
	err := i18n.ParseTranslationRoot("resource/translation")
	if err != nil {
		fmt.Printf("Could not install translations: %+v\n", err)
		return
	}
	r := util.Router(rootPath, i18n)
	http.ListenAndServe(":8080", r)
}
