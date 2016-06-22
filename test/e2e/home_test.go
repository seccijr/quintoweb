package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/seccijr/quintoweb/util"
	"github.com/seccijr/quintoweb/service"
	"os"
)

var (
	server   *httptest.Server
)

func init() {
	rootPath := os.Getenv("QUINTO_PATH");
	if rootPath  == "" {
		rootPath  = "/etc/root"
	}
	i18n := service.NewJsonI18n()
	i18n.ParseTranslationRoot("resource/translation")
	server = httptest.NewServer(util.Router(rootPath, i18n))
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
