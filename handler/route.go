package handler

import (
	"net/http"
)

func RouteInstall() {
	http.HandleFunc("/", Index)
}
