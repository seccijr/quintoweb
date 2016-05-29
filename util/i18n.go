package util

import (
	"golang.org/x/text/language"
	"encoding/json"
	"errors"
	"path/filepath"
	"os"
	"io/ioutil"
)

type TranslationMap map[string]string

var translations map[string]TranslationMap

func InitTranslations(supported []language.Tag) {
	translations = make(map[string]TranslationMap)
	for _, tag := range supported {
		translations[tag.String()] = make(TranslationMap)
	}
}

func AddTranslation(key string, trans string, tag language.Tag) error {
	if val, ok := translations[tag.String()]; ok {
		val[key] = trans
		return nil
	}
	return errors.New("No such language tag")
}

func GetTranslation(key string, tag language.Tag) (string, error) {
	if val, ok := translations[tag.String()]; ok {
		return val[key], nil
	}
	return "", errors.New("No such language tag")
}

func AddTranslationMap(languageMap TranslationMap, tag language.Tag) error {
	parent := tag.String()
	if _, ok := translations[parent]; ok {
		translations[parent] = languageMap
		return nil
	}
	return errors.New("No such language tag")
}

func ParseTranslations(raw []byte, tag language.Tag) error {
	languageMap := make(TranslationMap)
	json.Unmarshal(raw, &languageMap)
	return AddTranslationMap(languageMap, tag)
}

func parseTranslationFile(fullpath string, f os.FileInfo, err error) error {
	if (!f.IsDir()) {
		rawTag := filepath.Base(filepath.Dir(fullpath))
		tag := language.MustParse(rawTag)
		file, e := ioutil.ReadFile(fullpath)
		if e != nil {
			return e
		}
		return ParseTranslations(file, tag)
	}
	return nil
}

func ParseTranslationDir(path string) error {
	return filepath.Walk(path, parseTranslationFile)
}
