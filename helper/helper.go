package helper

import (
	"github.com/sailwith/goshopify/graphql"
	"github.com/sailwith/goshopify/rest"
)

type Helper struct {
	rest    *rest.REST
	graphQL *graphql.GraphQL
}

func NewHelper(
	rest *rest.REST,
	graphQL *graphql.GraphQL,
) *Helper {
	return &Helper{
		rest:    rest,
		graphQL: graphQL,
	}
}
