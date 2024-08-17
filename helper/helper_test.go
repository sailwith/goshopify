package helper

import (
	"testing"

	"github.com/sailwith/goshopify/graphql"
	"github.com/sailwith/goshopify/rest"
	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func newHelper(t *testing.T) *Helper {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := rest.NewREST(client)
	graphql := graphql.NewGraphQL(client)

	return NewHelper(rest, graphql)
}
