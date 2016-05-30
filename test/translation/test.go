package main

import (
	"github.com/seccijr/quintoweb/util"
	"golang.org/x/text/language"
	"fmt"
)

func main() {
	enTag := language.MustParse("en")
	esTag := language.MustParse("es")
	var supported = []language.Tag{enTag, esTag}
	util.InitTranslations(supported)
	err := util.ParseTranslationDir("resource/translation")
	if err != nil {
		fmt.Printf("Translation not found: %+v\n", err)
		return
	}
	hello, err := util.GetTranslation("hello", esTag)
	if err == nil {
		fmt.Printf("Translation from files hello: %s\n", hello)
	} else {
		fmt.Printf("Translation not found: %+v\n", err)
	}
}
