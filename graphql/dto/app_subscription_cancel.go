package dto

import "github.com/sailwith/goshopify/graphql/types"

const AppSubscriptionCancelQuery = `
mutation appSubscriptionCancel($id: ID!) {
  appSubscriptionCancel(id: $id) {
    appSubscription {
      id
    }
    userErrors {
      field
      message
    }
  }
}`

type AppSubscriptionCancelVars struct {
	ID string `json:"id"`
}

type AppSubscriptionCancelResp struct {
	AppSubscriptionCancel struct {
		AppSubscription types.AppSubscription `json:"appSubscription"`
		UserErrors      []types.UserError     `json:"userErrors"`
	} `json:"appSubscriptionCancel"`
}
