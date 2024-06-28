package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
)

func TestIsAppEmbedded(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	embedded, err := rest.IsAppEmbedded(context.Background(), test.BlockID)
	if err != nil {
		t.Error(err)
	}
	t.Log(embedded)
}
