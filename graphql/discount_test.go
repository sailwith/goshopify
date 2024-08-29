package graphql

import (
	"context"
	"testing"
	"time"

	"github.com/sailwith/goshopify/graphql/dto"
	"github.com/sailwith/goshopify/graphql/types"
	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestDiscountCodeBasicCreate(t *testing.T) {
	code := test.RandString(8)
	vars := dto.DiscountCodeBasicCreateVars{}
	vars.BasicCodeDiscount.Title = code
	vars.BasicCodeDiscount.CustomerGets.Value.Percentage = 0.1
	vars.BasicCodeDiscount.CustomerGets.Items.All = true
	vars.BasicCodeDiscount.UsageLimit = 1
	vars.BasicCodeDiscount.CustomerSelection.All = true

	startsAt := time.Now()
	endsAt := startsAt.Add(24 * time.Hour)
	vars.BasicCodeDiscount.StartsAt = time.Now()
	vars.BasicCodeDiscount.EndsAt = &endsAt
	vars.BasicCodeDiscount.Code = code
	vars.BasicCodeDiscount.AppliesOncePerCustomer = true

	graphQL := newGraphQL(t)
	resp, err := graphQL.DiscountCodeBasicCreate(context.Background(), vars)
	if assert.NoError(t, err) {
		t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID, code)
	}

	code = test.RandString(8)
	vars = dto.DiscountCodeBasicCreateVars{}
	vars.BasicCodeDiscount.Code = code
	vars.BasicCodeDiscount.CustomerGets.Items.All = true
	vars.BasicCodeDiscount.CustomerGets.Value.DiscountAmount = &types.DiscountAmountInput{
		Amount:            "1",
		AppliesOnEachItem: false,
	}
	vars.BasicCodeDiscount.CustomerSelection.All = true
	vars.BasicCodeDiscount.StartsAt = startsAt
	vars.BasicCodeDiscount.Title = code
	vars.BasicCodeDiscount.AppliesOncePerCustomer = true
	vars.BasicCodeDiscount.UsageLimit = 1
	resp, err = graphQL.DiscountCodeBasicCreate(context.Background(), vars)
	if assert.NoError(t, err) {
		t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID, code)
	}
}
