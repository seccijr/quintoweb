package repository

import (
	"database/sql"
	"github.com/seccijr/quintoweb/model"
	"golang.org/x/text/language"
	"time"
)

type Ad interface {
	GetTopDescOrder(limit int) []model.Ad
}

type AdPg struct {
	connection sql.DB
	lang       language.Tag
}

func NewAdPg(connection sql.DB, lang language.Tag) Ad {
	return AdPg{connection, lang}
}

func (ad AdPg) GetTopDescOrder(limit int) []model.Ad {
	return []model.Ad{
		model.Ad{"Ad 1", "Description 1", "picture1.jpg", time.Now()},
		model.Ad{"Ad 2", "Description 2", "picture2.jpg", time.Now()},
	}
}
