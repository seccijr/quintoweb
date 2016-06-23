package service

import (
	"encoding/json"
	"errors"
	"golang.org/x/text/language"
	"io/ioutil"
	"path/filepath"
)

type I18n interface {
	ParseTranslationRoot(path string) error
	GetTranslation(key string, tag language.Tag) string
	TransTemplate(args ...interface{}) (string, error)
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
	err := json.Unmarshal(raw, &languageMap)
	if err != nil {
		return err
	}
	return i18n.addTranslationMap(languageMap, tag)
}

func (i18n JsonI18n) parseTranslationDir(lang language.Tag, path string) error {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".json" {
			filename := filepath.Join(path, f.Name())
			raw, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			err = i18n.parseTranslations(raw, lang)
			if err != nil {
				return err
			}
		}
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
func (i18n JsonI18n) ParseTranslationRoot(path string) error {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			lang, err := language.Parse(f.Name())
			if err == nil {
				dir := filepath.Join(path, f.Name())
				err = i18n.parseTranslationDir(lang, dir)
			} else {
				return err
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (i18n JsonI18n) TransTemplate(args ...interface{}) (string, error) {
	if length := len(args); length == 2 {
		tag, tagOk := args[0].(language.Tag)
		key, keyOk := args[1].(string)
		if keyOk && tagOk {
			return i18n.GetTranslation(key, tag), nil
		} else {
			return "", errors.New("i18n: bad argument format")
		}
	}
	return "", errors.New("i18n: bad number of arguments")
}
