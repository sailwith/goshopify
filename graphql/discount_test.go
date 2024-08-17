package graphql

import (
	"context"
	"testing"
	"time"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestDiscountCodeBasicCreate(t *testing.T) {
	code := test.RandString(8)
	v := DiscountCodeBasicCreateVariable{}
	v.BasicCodeDiscount.Title = code
	v.BasicCodeDiscount.CustomerGets.Value.Percentage = 0.1
	v.BasicCodeDiscount.CustomerGets.Items.All = true
	v.BasicCodeDiscount.UsageLimit = 1
	v.BasicCodeDiscount.CustomerSelection.All = true

	startsAt := time.Now()
	endsAt := startsAt.Add(24 * time.Hour)
	v.BasicCodeDiscount.StartsAt = time.Now()
	v.BasicCodeDiscount.EndsAt = &endsAt
	v.BasicCodeDiscount.Code = code
	v.BasicCodeDiscount.AppliesOncePerCustomer = true

	graphQL := newGraphQL(t)
	resp, err := graphQL.DiscountCodeBasicCreate(context.Background(), v)
	if assert.NoError(t, err) {
		t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID)
	}

	code = test.RandString(8)
	v = DiscountCodeBasicCreateVariable{}
	v.BasicCodeDiscount.Code = code
	v.BasicCodeDiscount.CustomerGets.Items.All = true
	v.BasicCodeDiscount.CustomerGets.Value.DiscountAmount = &DiscountAmount{
		Amount:            1,
		AppliesOnEachItem: false,
	}
	v.BasicCodeDiscount.CustomerSelection.All = true
	v.BasicCodeDiscount.StartsAt = startsAt
	v.BasicCodeDiscount.Title = code
	v.BasicCodeDiscount.AppliesOncePerCustomer = true
	v.BasicCodeDiscount.UsageLimit = 1
	resp, err = graphQL.DiscountCodeBasicCreate(context.Background(), v)
	if assert.NoError(t, err) {
		t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID)
	}
}
