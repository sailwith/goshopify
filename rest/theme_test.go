package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThemeList(t *testing.T) {
	rest := newREST(t)
	themes, err := rest.ThemeList(context.Background())
	if assert.NoError(t, err) {
		for _, v := range themes {
			t.Logf("theme id: %v, name: %v, role: %v", v.Id, v.Name, v.Role)
		}
	}
}

func TestThemeMain(t *testing.T) {
	rest := newREST(t)
	theme, err := rest.ThemeMain(context.Background())
	if assert.NoError(t, err) {
		t.Logf("theme id: %v, name: %v, role: %v", theme.Id, theme.Name, theme.Role)
	}
}
