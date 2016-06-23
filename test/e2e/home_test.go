package e2e

import (
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/service"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

var (
	server *httptest.Server
)

func init() {
	root := environment.Root()
	i18n := service.NewJsonI18n()
	i18n.ParseTranslationRoot(filepath.Join(root, "resource/translation"))
	server = httptest.NewServer(handler.Router(i18n))
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
