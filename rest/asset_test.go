package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetGet(t *testing.T) {
	rest := newREST(t)
	mainTheme, err := rest.Theme.Main(context.Background())
	assert.NoError(t, err)

	asset, err := rest.Asset.Get(context.Background(), mainTheme.Id, "config/settings_data.json")
	if assert.NoError(t, err) {
		t.Log(asset.Key)
	}
}
