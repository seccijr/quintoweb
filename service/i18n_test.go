package service

import (
	"testing"
	"strings"
)

func TestGetExistingTranslation(t *testing.T) {
	// Arrange
	existingValue := "Existing Value"
	existingKey := "existingKey"
	language := "es"
	translationMap := make(map[string]TranslationMap)
	translationMap[language] = make(TranslationMap)
	translationMap[language][existingKey] = existingValue
	i18n := JsonI18n{translationMap}

	// Act
	value := i18n.GetTranslation(existingKey, language)

	// Assert
	if equal := strings.Compare(value, existingValue); equal != 0 {
		t.Error("Not equals values")
	}
}
