package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
)

func TestAssetGet(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	mainTheme, err := rest.ThemeMain(context.Background())
	if err != nil {
		t.Error(err)
	}

	asset, err := rest.AssetGet(context.Background(), mainTheme.Id, "config/settings_data.json")
	if err != nil {
		t.Error(err)
	}
	t.Log(asset.Key)
}
