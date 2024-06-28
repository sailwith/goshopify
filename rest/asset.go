package rest

import (
	"context"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Asset = boldshopify.Asset

func (r *REST) AssetGet(ctx context.Context, themeID uint64, key string) (*Asset, error) {
	return r.client.Asset.Get(ctx, themeID, key)
}
