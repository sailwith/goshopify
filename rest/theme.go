package rest

import (
	"context"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Theme = boldshopify.Theme

type ThemeResource struct {
	client *boldshopify.Client
}

func NewThemeResource(client *boldshopify.Client) *ThemeResource {
	return &ThemeResource{
		client: client,
	}
}

func (r *ThemeResource) List(ctx context.Context) ([]Theme, error) {
	return r.client.Theme.List(ctx, nil)
}
