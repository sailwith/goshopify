package test

import boldshopify "github.com/bold-commerce/go-shopify/v4"

const (
	ApiKey          = ""
	ApiSecret       = ""
	RedirectURL     = "http://127.0.0.1:8000/auth/callback"
	Scope           = "write_customers,write_discounts,read_themes"
	MyshopifyDomain = "foo.myshopify.com"
	ShopToken       = ""
	BlockID         = ""
	Version         = "2024-01"
)

func NewApp() *boldshopify.App {
	return &boldshopify.App{
		ApiKey:      ApiKey,
		ApiSecret:   ApiSecret,
		RedirectUrl: RedirectURL,
		Scope:       Scope,
	}
}

func NewClient() *boldshopify.Client {
	app := NewApp()
	return boldshopify.MustNewClient(
		*app,
		MyshopifyDomain,
		ShopToken,
		boldshopify.WithLogger(&boldshopify.LeveledLogger{
			Level: boldshopify.LevelDebug,
		}),
		boldshopify.WithVersion(Version),
	)
}
