package graphql

import (
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func newGraphQL(t *testing.T) *GraphQL {
	client, err := test.NewClient()
	assert.NoError(t, err)

	return NewGraphQL(client)
}
