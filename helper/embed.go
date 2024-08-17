package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type settingsData struct {
	Current struct {
		Blocks map[string]struct {
			Type     string `json:"type"`
			Disabled bool   `json:"disabled"`
		} `json:"blocks"`
	} `json:"current"`
}

func (h *Helper) IsAppEmbedded(ctx context.Context, blockID string) (bool, error) {
	mainTheme, err := h.rest.ThemeMain(ctx)
	if err != nil {
		return false, fmt.Errorf("error getting main theme: %w", err)
	}

	asset, err := h.rest.AssetGet(ctx, mainTheme.Id, "config/settings_data.json")
	if err != nil {
		return false, fmt.Errorf("error getting asset: %w", err)
	}

	var settingsData settingsData
	err = json.Unmarshal([]byte(asset.Value), &settingsData)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling asset value: %w [%s]", err, asset.Value)
	}

	if settingsData.Current.Blocks[blockID].Type == "" {
		return false, errors.New("block not found")
	}

	return !settingsData.Current.Blocks[blockID].Disabled, nil
}
