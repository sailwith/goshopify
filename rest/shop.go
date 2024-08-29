package rest

import (
	"context"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Shop = boldshopify.Shop

type ShopResource struct {
	client *boldshopify.Client
}

func NewShopResource(client *boldshopify.Client) *ShopResource {
	return &ShopResource{
		client: client,
	}
}

func (r *ShopResource) Get(ctx context.Context) (*Shop, error) {
	return r.client.Shop.Get(ctx, nil)
}
