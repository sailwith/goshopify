package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestThemeList(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	themes, err := rest.ThemeList(context.Background())
	if assert.NoError(t, err) {
		for _, v := range themes {
			t.Logf("theme id: %v, name: %v, role: %v", v.Id, v.Name, v.Role)
		}
	}
}

func TestThemeMain(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	theme, err := rest.ThemeMain(context.Background())
	if assert.NoError(t, err) {
		t.Logf("theme id: %v, name: %v, role: %v", theme.Id, theme.Name, theme.Role)
	}
}
