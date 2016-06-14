package service

import (
	"testing"
	"strings"
	"golang.org/x/text/language"
	"github.com/seccijr/quintoweb/service"
)

func TestGetExistingTranslation(t *testing.T) {
	// Arrange
	existingValue := "Existing Value"
	existingKey := "existingKey"
	rawTag := "es"
	tag := language.MustParse(rawTag)
	translations := make(map[string]service.TranslationMap)
	translations[rawTag] = make(service.TranslationMap)
	translations[rawTag][existingKey] = existingValue
	i18n := service.NewJsonI18nFeeded(translations)

	// Act
	value := i18n.GetTranslation(existingKey, tag)

	// Assert
	if equal := strings.Compare(value, existingValue); equal != 0 {
		t.Error("Not equals values")
	}
}

func TestGetNonExistingTranslation(t *testing.T) {
	// Arrange
	nonExistingKey := "nonexistingKey"
	existingValue := "Existing Value"
	existingKey := "existingKey"
	rawTag := "es"
	tag := language.MustParse(rawTag)
	translations := make(map[string]service.TranslationMap)
	translations[rawTag] = make(service.TranslationMap)
	translations[rawTag][existingKey] = existingValue
	i18n := service.NewJsonI18nFeeded(translations)

	// Act
	value := i18n.GetTranslation(nonExistingKey, tag)

	// Assert
	if equal := strings.Compare(value, existingValue); equal == 0 {
		t.Error("Retrieving non existing key")
	}
}

func TestGetNonExistingTranslationReturnsKey(t *testing.T) {
	// Arrange
	nonExistingKey := "nonexistingKey"
	existingValue := "Existing Value"
	existingKey := "existingKey"
	rawTag := "es"
	tag := language.MustParse(rawTag)
	translations := make(map[string]service.TranslationMap)
	translations[rawTag] = make(service.TranslationMap)
	translations[rawTag][existingKey] = existingValue
	i18n := service.NewJsonI18nFeeded(translations)

	// Act
	value := i18n.GetTranslation(nonExistingKey, tag)

	// Assert
	if equal := strings.Compare(value, nonExistingKey); equal != 0 {
		t.Error("The returned value should be equal to the non existing key")
	}
}

func TestParseProjectJsonDir(t *testing.T) {
	// Arrange
	i18n := service.NewJsonI18n()

	// Act
	err := i18n.ParseTranslationDir("../resource/translation")

	// Assert
	if err != nil {
		t.Error("Project translations are not in order")
	}
}

func XTestParseNonTranslationDir(t *testing.T) {
	// Arrange
	i18n := service.NewJsonI18n()

	// Act
	err := i18n.ParseTranslationDir("../resource/nontranslation")

	// Assert
	if err == nil {
		t.Error("Testing dir show fail")
	}
}
