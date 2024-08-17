package goshopify

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
	"github.com/sailwith/goshopify/graphql"
	"github.com/sailwith/goshopify/helper"
	"github.com/sailwith/goshopify/rest"
)

type App struct {
	instance boldshopify.App
	Version  string
}

type AppConfig struct {
	APIKey      string
	APISecret   string
	RedirectURL string
	Scope       string
	Version     string
}

type Client struct {
	REST    *rest.REST
	GraphQL *graphql.GraphQL
	Helper  *helper.Helper
}

func NewApp(cfg AppConfig) *App {
	instance := boldshopify.App{
		ApiKey:      cfg.APIKey,
		ApiSecret:   cfg.APISecret,
		RedirectUrl: cfg.RedirectURL,
		Scope:       cfg.Scope,
	}

	return &App{
		instance: instance,
		Version:  cfg.Version,
	}
}

func (app *App) NewClient(shopName string, token string) (*Client, error) {
	instance, err := boldshopify.NewClient(
		app.instance,
		shopName,
		token,
		boldshopify.WithRetry(3),
		boldshopify.WithLogger(&boldshopify.LeveledLogger{
			Level: boldshopify.LevelInfo,
		}),
		boldshopify.WithVersion(app.Version),
	)
	if err != nil {
		return nil, err
	}

	r := rest.NewREST(instance)
	g := graphql.NewGraphQL(instance)

	return &Client{
		REST:    r,
		GraphQL: g,
		Helper:  helper.NewHelper(r, g),
	}, nil
}

// AuthorizeURL generates the authorization URL for the Shopify app.
// e.g. https://example.myshopify.com/admin/oauth/authorize?client_id=apikey&redirect_uri=http%3A%2F%2F127.0.0.1%3A8000%2Fauth%2Fcallback&scope=write_customers%2Cwrite_discounts%2Cread_themes&state=state
func (app *App) AuthorizeURL(myshopifyDomain string, state string) (string, error) {
	return app.instance.AuthorizeUrl(myshopifyDomain, state)
}

// EmbeddedAppURL generates the embedded app URL for the Shopify app.
// e.g. https://admin.shopify.com/store/example/apps/apikey
func (app *App) EmbeddedAppURL(host string) (string, error) {
	decodeHost, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(host)
	if err != nil {
		return "", fmt.Errorf("error decoding host: %v", err)
	}

	return "https://" + string(decodeHost) + "/apps/" + app.instance.ApiKey, nil
}

// VerifyAuthorizationURL verifies that the OAuth callback request is from Shopify.
func (app *App) VerifyAuthorizationURL(u *url.URL) (bool, error) {
	return app.instance.VerifyAuthorizationURL(u)
}

// GetAccessToken exchanges the authorization code from the confirmation redirect for an access token.
func (app *App) GetAccessToken(ctx context.Context, shopName string, code string) (string, error) {
	return app.instance.GetAccessToken(ctx, shopName, code)
}

func (app *App) VerifyWebhookRequest(req *http.Request) bool {
	return app.instance.VerifyWebhookRequest(req)
}
