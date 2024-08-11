package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestIsAppEmbedded(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	embedded, err := rest.IsAppEmbedded(context.Background(), test.BlockID)
	if assert.NoError(t, err) {
		t.Log(embedded)
	}
}
