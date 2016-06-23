package model

import (
	"golang.org/x/text/language"
	"time"
)

type Base struct {
	Title string
	Body  []byte
	Lang  language.Tag
}

type Ad struct {
	Title       string
	Description string
	Picture     string
	Deadline    time.Time
}

type Index struct {
	Base
	Ads []Ad
}
