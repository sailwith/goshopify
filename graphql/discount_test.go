package graphql

import (
	"context"
	"testing"
	"time"

	"github.com/sailwith/goshopify/test"
)

func TestDiscountCodeBasicCreate(t *testing.T) {
	v := DiscountCodeBasicCreateVariable{}
	v.BasicCodeDiscount.Title = "123"
	v.BasicCodeDiscount.CustomerGets.Value.Percentage = 0.1
	v.BasicCodeDiscount.CustomerGets.Items.All = true
	v.BasicCodeDiscount.UsageLimit = 1
	v.BasicCodeDiscount.CustomerSelection.All = true
	v.BasicCodeDiscount.StartsAt = time.Now()
	v.BasicCodeDiscount.Code = "11"
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
}
