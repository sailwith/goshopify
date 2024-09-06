package helper

import (
	"context"
	"fmt"

	"github.com/sailwith/goshopify/rest"
)

func (h *Helper) CreateOrUpdateCustomer(ctx context.Context, customer rest.Customer) (*rest.Customer, error) {
	customers, err := h.rest.Customer.Search(ctx, "email:"+customer.Email)
	if err != nil {
		return nil, fmt.Errorf("error searching for customer: %w", err)
	}

	shopifyCustomer := new(rest.Customer)
	if len(customers) == 0 {
		shopifyCustomer, err = h.rest.Customer.Create(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error creating customer: %w", err)
		}
	} else {
		customer.Id = customers[0].Id
		shopifyCustomer, err = h.rest.Customer.Update(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error updating customer: %w", err)
		}
	}

	return shopifyCustomer, nil
}
