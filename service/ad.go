package service

import (
	"github.com/seccijr/quintoweb/model"
	"github.com/seccijr/quintoweb/repository"
	"golang.org/x/text/language"
)

type Ad interface {
	GetTopDescOrder(limit int) []model.Ad
}

type AdI15d struct {
	lang         language.Tag
	adRepository repository.Ad
}

func NewAdI15d(adRepository repository.Ad, lang language.Tag) Ad {
	return AdI15d{lang, adRepository}
}

func (ad AdI15d) GetTopDescOrder(limit int) []model.Ad {
	return ad.adRepository.GetTopDescOrder(limit)
}
