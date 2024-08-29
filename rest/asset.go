package rest

import (
	"context"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Asset = boldshopify.Asset

type AssetResource struct {
	client *boldshopify.Client
}

func NewAssetResource(client *boldshopify.Client) *AssetResource {
	return &AssetResource{
		client: client,
	}
}

func (r *AssetResource) Get(ctx context.Context, themeID uint64, key string) (*Asset, error) {
	return r.client.Asset.Get(ctx, themeID, key)
}
