package helper

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestIsAppEmbedded(t *testing.T) {
	h := newHelper(t)
	embedded, err := h.IsAppEmbedded(context.Background(), test.BlockID)
	if assert.NoError(t, err) {
		t.Log(embedded)
	}
}
