package dto

import "github.com/sailwith/goshopify/graphql/types"

const AppSubscriptionCreateQuery = `
mutation appSubscriptionCreate($lineItems: [AppSubscriptionLineItemInput!]!, $name: String!, $returnUrl: URL!, $trialDays: Int, $test: Boolean) {
  appSubscriptionCreate(lineItems: $lineItems, name: $name, returnUrl: $returnUrl, trialDays: $trialDays, test: $test) {
    appSubscription {
      id
      returnUrl
    }
    confirmationUrl
    userErrors {
      field
      message
    }
  }
}`

type AppSubscriptionCreateVars struct {
	LineItems []types.AppSubscriptionLineItemInput `json:"lineItems"`
	Name      string                               `json:"name"`
	ReturnURL string                               `json:"returnUrl"`
	TrialDays int                                  `json:"trialDays"`
	Test      bool                                 `json:"test"`
}

type AppSubscriptionCreateResp struct {
	AppSubscriptionCreate struct {
		AppSubscription types.AppSubscription `json:"appSubscription"`
		ConfirmationUrl string                `json:"confirmationUrl"`
		UserErrors      []types.UserError     `json:"userErrors"`
	} `json:"appSubscriptionCreate"`
}
