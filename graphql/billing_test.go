package graphql

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/graphql/dto"
	"github.com/sailwith/goshopify/graphql/types"
	"github.com/stretchr/testify/assert"
)

func TestAppSubscriptionCreate(t *testing.T) {
	graphQL := newGraphQL(t)

	vars := dto.AppSubscriptionCreateVars{
		LineItems: []types.AppSubscriptionLineItemInput{
			{
				Plan: types.AppPlanInput{
					AppRecurringPricingDetails: &types.AppRecurringPricingInput{
						Discount: &types.AppSubscriptionDiscountInput{
							DurationLimitInIntervals: 2,
							Value: types.AppSubscriptionDiscountValueInput{
								Amount: "5",
							},
						},
						Interval: "EVERY_30_DAYS",
						Price: types.MoneyInput{
							Amount:       "29.9",
							CurrencyCode: "USD",
						},
					},
				},
			},
			{
				Plan: types.AppPlanInput{
					AppUsagePricingDetails: &types.AppUsagePricingInput{
						CappedAmount: types.MoneyInput{
							Amount:       "999",
							CurrencyCode: "USD",
						},
						Terms: "message",
					},
				},
			},
		},
		Name:      "Plan",
		ReturnURL: "https://example.com",
		TrialDays: 3,
		Test:      true,
	}
	resp, err := graphQL.AppSubscriptionCreate(context.Background(), vars)
	if assert.NoError(t, err) {
		t.Log(resp)
	}
}
