package dto

import "github.com/sailwith/goshopify/graphql/types"

const DiscountCodeBasicCreateQuery = `
mutation discountCodeBasicCreate($basicCodeDiscount: DiscountCodeBasicInput!) {
  discountCodeBasicCreate(basicCodeDiscount: $basicCodeDiscount) {
    codeDiscountNode {
      id
    }
    userErrors {
	  code
	  extraInfo
      field
      message
    }
  }
}`

type DiscountCodeBasicCreateVars struct {
	BasicCodeDiscount types.DiscountCodeBasicInput `json:"basicCodeDiscount"`
}

type DiscountCodeBasicCreateResp struct {
	DiscountCodeBasicCreate struct {
		CodeDiscountNode types.DiscountCodeNode    `json:"codeDiscountNode"`
		UserErrors       []types.DiscountUserError `json:"userErrors"`
	} `json:"discountCodeBasicCreate"`
}
