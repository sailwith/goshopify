package graphql

import (
	"context"
	"fmt"

	"github.com/sailwith/goshopify/graphql/dto"
)

func (g *GraphQL) DiscountCodeBasicCreate(ctx context.Context, vars dto.DiscountCodeBasicCreateVars) (*dto.DiscountCodeBasicCreateResp, error) {
	resp := dto.DiscountCodeBasicCreateResp{}
	if err := g.client.GraphQL.Query(ctx, dto.DiscountCodeBasicCreateQuery, vars, &resp); err != nil {
		return nil, err
	}
	if len(resp.DiscountCodeBasicCreate.UserErrors) > 0 {
		return nil, fmt.Errorf("%v", resp.DiscountCodeBasicCreate.UserErrors)
	}

	return &resp, nil
}
