package types

import "time"

type AppPlanInput struct {
	AppRecurringPricingDetails *AppRecurringPricingInput `json:"appRecurringPricingDetails,omitempty"`
	AppUsagePricingDetails     *AppUsagePricingInput     `json:"appUsagePricingDetails,omitempty"`
}

type AppRecurringPricingInput struct {
	Discount *AppSubscriptionDiscountInput `json:"discount,omitempty"`
	Interval string                        `json:"interval"`
	Price    MoneyInput                    `json:"price"`
}

type AppSubscriptionDiscountInput struct {
	DurationLimitInIntervals int                               `json:"durationLimitInIntervals"`
	Value                    AppSubscriptionDiscountValueInput `json:"value"`
}

type AppSubscriptionDiscountValueInput struct {
	Amount     string  `json:"amount,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

type AppSubscriptionLineItemInput struct {
	Plan AppPlanInput `json:"plan"`
}

type AppUsagePricingInput struct {
	CappedAmount MoneyInput `json:"cappedAmount"`
	Terms        string     `json:"terms"`
}

type DiscountAmountInput struct {
	Amount            string `json:"amount"`
	AppliesOnEachItem bool   `json:"appliesOnEachItem"`
}

type DiscountCodeBasicInput struct {
	AppliesOncePerCustomer bool                           `json:"appliesOncePerCustomer"`
	Code                   string                         `json:"code"`
	CustomerGets           DiscountCustomerGetsInput      `json:"customerGets"`
	CustomerSelection      DiscountCustomerSelectionInput `json:"customerSelection"`
	EndsAt                 *time.Time                     `json:"endsAt"`
	StartsAt               time.Time                      `json:"startsAt"`
	Title                  string                         `json:"title"`
	UsageLimit             int                            `json:"usageLimit"`
}

type DiscountCustomerGetsInput struct {
	Items DiscountItemsInput             `json:"items"`
	Value DiscountCustomerGetsValueInput `json:"value"`
}

type DiscountCustomerGetsValueInput struct {
	DiscountAmount *DiscountAmountInput `json:"discountAmount,omitempty"`
	Percentage     float64              `json:"percentage,omitempty"`
}

type DiscountCustomerSelectionInput struct {
	All bool `json:"all"`
}

type DiscountItemsInput struct {
	All bool `json:"all"`
}

type MoneyInput struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}
