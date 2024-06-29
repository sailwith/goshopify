package test

import (
	"math/rand"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

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

func NewClient() (*boldshopify.Client, error) {
	app := NewApp()
	return boldshopify.NewClient(
		*app,
		MyshopifyDomain,
		ShopToken,
		boldshopify.WithLogger(&boldshopify.LeveledLogger{
			Level: boldshopify.LevelDebug,
		}),
		boldshopify.WithVersion(Version),
	)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
