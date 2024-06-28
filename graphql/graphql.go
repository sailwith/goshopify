package graphql

import boldshopify "github.com/bold-commerce/go-shopify/v4"

type GraphQL struct {
	client *boldshopify.Client
}

func NewGraphQL(client *boldshopify.Client) *GraphQL {
	return &GraphQL{
		client: client,
	}
}
