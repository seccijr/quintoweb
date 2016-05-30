package util

import (
	"golang.org/x/text/language"
	"encoding/json"
	"errors"
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
)

type TranslationMap map[string]string

var translations map[string]TranslationMap = make(map[string]TranslationMap)

func addTranslationMap(languageMap TranslationMap, tag language.Tag) error {
	parent := tag.String()
	translations[parent] = languageMap
	return nil
}

func parseTranslations(raw []byte, tag language.Tag) error {
	languageMap := make(TranslationMap)
	json.Unmarshal(raw, &languageMap)
	return addTranslationMap(languageMap, tag)
}

func parseTranslationFile(fullpath string, f os.FileInfo, err error) error {
	if (!f.IsDir()) {
		rawTag := filepath.Base(filepath.Dir(fullpath))
		tag := language.MustParse(rawTag)
		file, e := ioutil.ReadFile(fullpath)
		if e != nil {
			return e
		}
		return parseTranslations(file, tag)
	}
	return nil
}

// Walks through a translation directory looking for
// json translation files named as LANG.json where
// LANG must be replaced with the language code corresponding
// to the translation file
func ParseTranslationDir(path string) error {
	return filepath.Walk(path, parseTranslationFile)
}

// Gets the translation for the specified key in the language
// represented by tag
func GetTranslation(key string, tag language.Tag) (string, error) {
	if val, ok := translations[tag.String()]; ok {
		return val[key], nil
	}
	return "", errors.New("No such language tag")
}

// Aux function used in order to translate templates
func TransTemplate(args... interface{}) (string, error) {
	if length := len(args); length == 2 {
		fmt.Printf("Called with: \n%+v\n%+v\n", args[0], args[1])
		tag, tagOk := args[0].(language.Tag)
		key, keyOk := args[1].(string)
		if keyOk && tagOk {
			return GetTranslation(key, tag)
		} else {
			return "", errors.New("Bad argument format")
		}
	}
	return "", errors.New("Not enough arguments")
}
