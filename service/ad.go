package service

import (
	"github.com/seccijr/quintoweb/model"
	"time"
)

func GetTopAdsDescOrder(limit int) []model.Ad {
	return []model.Ad{
		model.Ad{"Ad 1", "Description 1", "picture1.jpg", time.Now()},
		model.Ad{"Ad 2", "Description 2", "picture2.jpg", time.Now()},
	}
}
