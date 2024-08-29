package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShopGet(t *testing.T) {
	rest := newREST(t)
	shop, err := rest.Shop.Get(context.Background())
	if assert.NoError(t, err) {
		t.Log(shop.MyshopifyDomain)
	}
}
