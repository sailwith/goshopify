package rest

import (
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func newREST(t *testing.T) *REST {
	client, err := test.NewClient()
	assert.NoError(t, err)

	return NewREST(client)
}
