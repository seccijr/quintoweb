package server

import (
	"github.com/seccijr/quintoweb/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
)

func init() {
	mockI18n := MockI18n{}
	mockAdService := MockAdService{}
	server = httptest.NewServer(handler.Router(mockI18n, mockAdService))
}

func TestRespondingHomePage(t *testing.T) {
	res, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
