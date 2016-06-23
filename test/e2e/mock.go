package e2e

import (
	"github.com/seccijr/quintoweb/model"
	"time"
	"golang.org/x/text/language"
	"errors"
)


type MockI18n struct {
}

func (i18n MockI18n) ParseTranslationRoot(path string) error {
	return nil
}

func (i18n MockI18n) GetTranslation(key string, tag language.Tag) string {
	return key
}

func (i18n MockI18n) TransTemplate(args ...interface{}) (string, error) {
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

type MockAdService struct {
}

func (ad MockAdService) GetTopDescOrder(limit int) []model.Ad {
	return []model.Ad{
		model.Ad{"Ad 1", "Description 1", "picture1.jpg", time.Now()},
		model.Ad{"Ad 2", "Description 2", "picture2.jpg", time.Now()},
	}
}
