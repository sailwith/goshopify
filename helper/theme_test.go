package helper

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainTheme(t *testing.T) {
	h := newHelper(t)
	theme, err := h.MainTheme(context.Background())
	if assert.NoError(t, err) {
		t.Logf("theme id: %v, name: %v, role: %v", theme.Id, theme.Name, theme.Role)
	}
}
