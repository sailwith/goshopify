package graphql

import (
	"context"
	"time"
)

const q = `
mutation discountCodeBasicCreate($basicCodeDiscount: DiscountCodeBasicInput!) {
  discountCodeBasicCreate(basicCodeDiscount: $basicCodeDiscount) {
    codeDiscountNode {
      id
    }
    userErrors {
      field
      code
      message
    }
  }
}`

type v struct {
	BasicCodeDiscount struct {
		Code         string `json:"code"`
		CustomerGets struct {
			Items struct {
				All bool `json:"all"`
			} `json:"items"`
			Value struct {
				DiscountAmount *struct {
					Amount            float64 `json:"amount"`
					AppliesOnEachItem bool    `json:"appliesOnEachItem"`
				} `json:"discountAmount,omitempty"`
				Percentage float64 `json:"percentage,omitempty"`
			} `json:"value"`
		} `json:"customerGets"`
		CustomerSelection struct {
			All bool `json:"all"`
		} `json:"customerSelection"`
		StartsAt               time.Time `json:"startsAt"`
		Title                  string    `json:"title"`
		AppliesOncePerCustomer bool      `json:"appliesOncePerCustomer"`
		UsageLimit             int       `json:"usageLimit"`
	} `json:"basicCodeDiscount"`
}

type resp struct {
	DiscountCodeBasicCreate struct {
		CodeDiscountNode struct {
			ID string `json:"id"`
		} `json:"codeDiscountNode"`
		UserErrors []struct {
			Field   []string `json:"field"`
			Code    string   `json:"code"`
			Message string   `json:"message"`
		} `json:"userErrors"`
	} `json:"discountCodeBasicCreate"`
}

func (g *GraphQL) DiscountCodeBasicCreate(ctx context.Context, variable v) (*resp, error) {
	resp := new(resp)
	if err := g.client.GraphQL.Query(ctx, q, variable, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
