package goshopify

import (
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func newApp() *App {
	return NewApp(AppConfig{
		APIKey:      test.ApiKey,
		APISecret:   test.ApiSecret,
		RedirectURL: test.RedirectURL,
		Scope:       test.Scope,
		Version:     test.Version,
	})
}

func TestAuthorizeURL(t *testing.T) {
	app := newApp()
	url, err := app.AuthorizeURL(test.MyshopifyDomain, "state")
	if assert.NoError(t, err) {
		t.Log(url)
	}
}

func TestEmbeddedAppURL(t *testing.T) {
	app := newApp()
	url, err := app.EmbeddedAppURL("YWRtaW4uc2hvcGlmeS5jb20vc3RvcmUvZXhhbXBsZQ")
	if assert.NoError(t, err) {
		t.Log(url)
	}
}
