package graphql

import (
	"context"
	"fmt"

	"github.com/sailwith/goshopify/graphql/dto"
)

func (g *GraphQL) AppSubscriptionCreate(ctx context.Context, vars dto.AppSubscriptionCreateVars) (*dto.AppSubscriptionCreateResp, error) {
	resp := dto.AppSubscriptionCreateResp{}
	if err := g.client.GraphQL.Query(ctx, dto.AppSubscriptionCreateQuery, vars, &resp); err != nil {
		return nil, err
	}
	if len(resp.AppSubscriptionCreate.UserErrors) > 0 {
		return nil, fmt.Errorf("%v", resp.AppSubscriptionCreate.UserErrors)
	}
	return &resp, nil
}

func (g *GraphQL) AppSubscriptionCancel(ctx context.Context, vars dto.AppSubscriptionCancelVars) (*dto.AppSubscriptionCancelResp, error) {
	resp := dto.AppSubscriptionCancelResp{}
	if err := g.client.GraphQL.Query(ctx, dto.AppSubscriptionCancelQuery, vars, &resp); err != nil {
		return nil, err
	}
	if len(resp.AppSubscriptionCancel.UserErrors) > 0 {
		return nil, fmt.Errorf("%v", resp.AppSubscriptionCancel.UserErrors)
	}
	return &resp, nil
}
