package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetGet(t *testing.T) {
	rest := newREST(t)
	themes, err := rest.Theme.List(context.Background())
	assert.NoError(t, err)

	asset, err := rest.Asset.Get(context.Background(), themes[0].Id, "config/settings_data.json")
	if assert.NoError(t, err) {
		t.Log(asset.Key)
	}
}
