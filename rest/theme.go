package rest

import (
	"context"
	"errors"
	"fmt"

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

func (r *ThemeResource) Main(ctx context.Context) (*Theme, error) {
	themes, err := r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting themes: %w", err)
	}

	var mainTheme Theme
	for _, v := range themes {
		if v.Role == "main" {
			mainTheme = v
		}
	}
	if mainTheme.Id == 0 {
		return nil, errors.New("main theme not found")
	}

	return &mainTheme, nil
}
