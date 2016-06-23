package service

import (
	"github.com/seccijr/quintoweb/service"
	"golang.org/x/text/language"
	"strings"
	"testing"
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
	tag := language.MustParse("es")

	// Act
	err := i18n.ParseTranslationRoot("../resource/translation")
	value := i18n.GetTranslation("test", tag)

	// Assert
	if err != nil {
		t.Error("Project translations are not in order")
	}
	if equal := strings.Compare(value, "Test"); equal != 0 {
		t.Error("The returned value should be equal to the Test value")
	}
}

func TestParseBadStructureTranslationDir(t *testing.T) {
	// Arrange
	i18n := service.NewJsonI18n()

	// Act
	err := i18n.ParseTranslationRoot("../resource/badstructuretranslation")

	// Assert
	if err == nil {
		t.Error("Testing dir may show some failure")
	}
	if err.Error() != "language: tag is not well-formed" {
		t.Error("Returned error may show language tag is not well-formed failure")
	}
}

func TestParseBadFormatTranslation(t *testing.T) {
	// Arrange
	i18n := service.NewJsonI18n()

	// Act
	err := i18n.ParseTranslationRoot("../resource/badformattranslation")

	// Assert
	if err == nil {
		t.Error("Testing dir may show some failure")
	}
	if !strings.Contains(err.Error(), "json: cannot unmarshal string into Go value") {
		t.Error("Returned error may show cannot unmarshal string into Go value failure")
	}
}

func TestTranslateExistingTemplateKey(t *testing.T) {
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
	value, err := i18n.TransTemplate(tag, existingKey)

	// Assert
	if err != nil {
		t.Errorf("Some error rised up: %v\n", err)
	}
	if equal := strings.Compare(value, existingValue); equal != 0 {
		t.Error("Not equals values")
	}
}

func TestBadNArgsCallToTranslateTemplate(t *testing.T) {
	// Arrange
	translations := make(map[string]service.TranslationMap)
	i18n := service.NewJsonI18nFeeded(translations)

	// Act
	_, err := i18n.TransTemplate("Arg 1", "Arg 2", "Arg 3")

	// Assert
	if err == nil {
		t.Error("Testing template may show some failure")
	}
	if err.Error() != "i18n: bad number of arguments" {
		t.Error("Returned error may i18n show bad number of arguments failure")
	}
}

func TestBadFormatArgsCallToTranslateTemplate(t *testing.T) {
	// Arrange
	translations := make(map[string]service.TranslationMap)
	i18n := service.NewJsonI18nFeeded(translations)

	// Act
	_, err := i18n.TransTemplate("Arg 1", "Arg 2")

	// Assert
	if err == nil {
		t.Error("Testing template may show some failure")
	}
	if err.Error() != "i18n: bad argument format" {
		t.Error("Returned error may i18n bad argument format failure")
	}
}
