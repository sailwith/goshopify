package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
)

func TestShopGet(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	shop, err := rest.ShopGet(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(shop.MyshopifyDomain)
}
