package rest

import boldshopify "github.com/bold-commerce/go-shopify/v4"

type REST struct {
	client *boldshopify.Client
}

func NewREST(client *boldshopify.Client) *REST {
	return &REST{
		client: client,
	}
}
