package helper

import (
	"context"
	"errors"
	"fmt"

	"github.com/sailwith/goshopify/rest"
)

func (h *Helper) MainTheme(ctx context.Context) (*rest.Theme, error) {
	themes, err := h.rest.Theme.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting themes: %w", err)
	}

	var mainTheme rest.Theme
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
