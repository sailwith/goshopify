package rest

import (
	"context"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Shop = boldshopify.Shop

func (r *REST) ShopGet(ctx context.Context) (*Shop, error) {
	return r.client.Shop.Get(ctx, nil)
}
