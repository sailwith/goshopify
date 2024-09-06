package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThemeList(t *testing.T) {
	rest := newREST(t)
	themes, err := rest.Theme.List(context.Background())
	if assert.NoError(t, err) {
		for _, v := range themes {
			t.Logf("theme id: %v, name: %v, role: %v", v.Id, v.Name, v.Role)
		}
	}
}
