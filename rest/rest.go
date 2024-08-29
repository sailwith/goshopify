package rest

import boldshopify "github.com/bold-commerce/go-shopify/v4"

type REST struct {
	Asset    *AssetResource
	Customer *CustomerResource
	Shop     *ShopResource
	Theme    *ThemeResource
}

func NewREST(client *boldshopify.Client) *REST {
	return &REST{
		Asset:    NewAssetResource(client),
		Customer: NewCustomerResource(client),
		Shop:     NewShopResource(client),
		Theme:    NewThemeResource(client),
	}
}
