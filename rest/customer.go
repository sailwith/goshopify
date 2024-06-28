package rest

import (
	"context"
	"fmt"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type Customer = boldshopify.Customer

func (r *REST) CustomerCreate(ctx context.Context, customer Customer) (*Customer, error) {
	return r.client.Customer.Create(ctx, customer)
}

func (r *REST) CustomerSearch(ctx context.Context, query string) ([]Customer, error) {
	options := boldshopify.CustomerSearchOptions{
		Query: query,
	}
	return r.client.Customer.Search(ctx, options)
}

func (r *REST) CustomerUpdate(ctx context.Context, customer Customer) (*Customer, error) {
	return r.client.Customer.Update(ctx, customer)
}

func (r *REST) CustomerCreateOrUpdate(ctx context.Context, customer Customer) (*Customer, error) {
	customers, err := r.CustomerSearch(ctx, "email:"+customer.Email)
	if err != nil {
		return nil, fmt.Errorf("error searching for customer: %w", err)
	}

	shopifyCustomer := new(Customer)
	if len(customers) == 0 {
		shopifyCustomer, err = r.CustomerCreate(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error creating customer: %w", err)
		}
	} else {
		customer.Id = customers[0].Id
		shopifyCustomer, err = r.CustomerUpdate(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error updating customer: %w", err)
		}
	}

	return shopifyCustomer, nil
}
