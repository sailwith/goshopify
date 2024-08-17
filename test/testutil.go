package test

import (
	"math/rand"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

const (
	ApiKey          = "61b03e068f12323d4973b9ff1b7148c4"
	ApiSecret       = "7bbe4b0f8310f1e409887ebfa827ff15"
	RedirectURL     = "http://127.0.0.1:8000/auth/callback"
	Scope           = "write_customers,write_discounts,read_themes"
	MyshopifyDomain = "spf-dev1.myshopify.com"
	ShopToken       = "shpua_5d3191367415e7206fcb41e83c71a70b"
	BlockID         = "1486664396617447524"
	Version         = "2024-07"
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
			Level: boldshopify.LevelInfo,
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
