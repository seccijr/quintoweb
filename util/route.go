package util

import (
	"net/http"
)

func RouteInstall() {
	http.HandleFunc("/", handler.Index)
}
