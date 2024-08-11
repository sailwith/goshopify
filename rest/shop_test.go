package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestShopGet(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	shop, err := rest.ShopGet(context.Background())
	if assert.NoError(t, err) {
		t.Log(shop.MyshopifyDomain)
	}
}
