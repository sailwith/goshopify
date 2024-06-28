package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
)

func TestThemeList(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	themes, err := rest.ThemeList(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, v := range themes {
		t.Logf("theme id: %v, name: %v, role: %v", v.Id, v.Name, v.Role)
	}
}

func TestThemeMain(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	theme, err := rest.ThemeMain(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Logf("theme id: %v, name: %v, role: %v", theme.Id, theme.Name, theme.Role)
}
