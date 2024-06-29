package graphql

import (
	"context"
	"testing"
	"time"

	"github.com/sailwith/goshopify/test"
)

func TestDiscountCodeBasicCreate(t *testing.T) {
	code := test.RandString(8)
	v := DiscountCodeBasicCreateVariable{}
	v.BasicCodeDiscount.Title = code
	v.BasicCodeDiscount.CustomerGets.Value.Percentage = 0.1
	v.BasicCodeDiscount.CustomerGets.Items.All = true
	v.BasicCodeDiscount.UsageLimit = 1
	v.BasicCodeDiscount.CustomerSelection.All = true
	v.BasicCodeDiscount.StartsAt = time.Now()
	v.BasicCodeDiscount.Code = code
	v.BasicCodeDiscount.AppliesOncePerCustomer = true
	client, err := test.NewClient()
	if err != nil {
		t.Error(err)
	}
	graphQL := NewGraphQL(client)
	resp, err := graphQL.DiscountCodeBasicCreate(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID)

	code = test.RandString(8)
	v = DiscountCodeBasicCreateVariable{}
	v.BasicCodeDiscount.Code = code
	v.BasicCodeDiscount.CustomerGets.Items.All = true
	v.BasicCodeDiscount.CustomerGets.Value.DiscountAmount = &DiscountAmount{
		Amount:            1,
		AppliesOnEachItem: false,
	}
	v.BasicCodeDiscount.CustomerSelection.All = true
	v.BasicCodeDiscount.StartsAt = time.Now()
	v.BasicCodeDiscount.Title = code
	v.BasicCodeDiscount.AppliesOncePerCustomer = true
	v.BasicCodeDiscount.UsageLimit = 1
	resp, err = graphQL.DiscountCodeBasicCreate(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.DiscountCodeBasicCreate.CodeDiscountNode.ID)
}
