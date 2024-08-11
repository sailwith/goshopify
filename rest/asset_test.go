package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestAssetGet(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	mainTheme, err := rest.ThemeMain(context.Background())
	assert.NoError(t, err)

	asset, err := rest.AssetGet(context.Background(), mainTheme.Id, "config/settings_data.json")
	if assert.NoError(t, err) {
		t.Log(asset.Key)
	}
}
