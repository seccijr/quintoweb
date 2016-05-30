package model

import (
	"golang.org/x/text/language"
)

type Page struct {
	Title string
	Body  []byte
	Lang  language.Tag
}
