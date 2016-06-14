package service

import (
	"golang.org/x/text/language"
	"encoding/json"
	"errors"
	"path/filepath"
	"os"
	"io/ioutil"
)

type I18n interface {
	ParseTranslationDir(path string) error
	GetTranslation(key string, tag language.Tag) string
	TransTemplate(args... interface{}) (string, error)
}

type TranslationMap map[string]string

type JsonI18n struct {
	translations map[string]TranslationMap
}

func (i18n JsonI18n) addTranslationMap(languageMap TranslationMap, tag language.Tag) error {
	parent := tag.String()
	i18n.translations[parent] = languageMap
	return nil
}

func (i18n JsonI18n) parseTranslations(raw []byte, tag language.Tag) error {
	languageMap := make(TranslationMap)
	json.Unmarshal(raw, &languageMap)
	return i18n.addTranslationMap(languageMap, tag)
}

func (i18n JsonI18n) parseTranslationFile(fullpath string, f os.FileInfo, err error) error {
	if (!f.IsDir()) {
		rawTag := filepath.Base(filepath.Dir(fullpath))
		tag := language.MustParse(rawTag)
		file, e := ioutil.ReadFile(fullpath)
		if e != nil {
			return e
		}
		return i18n.parseTranslations(file, tag)
	}
	return nil
}

// Creates a new instance of type JsonI18n
func NewJsonI18n() I18n {
	i18n := JsonI18n{}
	i18n.translations = make(map[string]TranslationMap)
	return i18n
}

// Creates a new instance of type JsonI18n
func NewJsonI18nFeeded(translations map[string]TranslationMap) I18n {
	i18n := JsonI18n{}
	i18n.translations = translations
	return i18n
}

// Walks through a translation directory looking for
// json translation files named as LANG.json where
// LANG must be replaced with the language code corresponding
// to the translation file
func (i18n JsonI18n) ParseTranslationDir(path string) error {
	return filepath.Walk(path, i18n.parseTranslationFile)
}

// Gets the translation for the specified key in the language
// represented by tag
func (i18n JsonI18n) GetTranslation(key string, tag language.Tag) string {
	if transMap, okMap := i18n.translations[tag.String()]; okMap {
		if val, okVal := transMap[key]; okVal {
			return val
		}
	}
	return key
}

// Aux function used in order to translate templates
func (i18n JsonI18n) TransTemplate(args... interface{}) (string, error) {
	if length := len(args); length == 2 {
		tag, tagOk := args[0].(language.Tag)
		key, keyOk := args[1].(string)
		if keyOk && tagOk {
			return i18n.GetTranslation(key, tag), nil
		} else {
			return "", errors.New("Bad argument format")
		}
	}
	return "", errors.New("Not enough arguments")
}
