package rest

import (
	"context"
	"fmt"

	boldshopify "github.com/bold-commerce/go-shopify/v4"
)

type EmailMarketingConsent = boldshopify.EmailMarketingConsent
type Customer = boldshopify.Customer

type CustomerResource struct {
	client *boldshopify.Client
}

func NewCustomerResource(client *boldshopify.Client) *CustomerResource {
	return &CustomerResource{
		client: client,
	}
}

func (r *CustomerResource) Create(ctx context.Context, customer Customer) (*Customer, error) {
	return r.client.Customer.Create(ctx, customer)
}

func (r *CustomerResource) Search(ctx context.Context, query string) ([]Customer, error) {
	options := boldshopify.CustomerSearchOptions{
		Query: query,
	}
	return r.client.Customer.Search(ctx, options)
}

func (r *CustomerResource) Update(ctx context.Context, customer Customer) (*Customer, error) {
	return r.client.Customer.Update(ctx, customer)
}

func (r *CustomerResource) CreateOrUpdate(ctx context.Context, customer Customer) (*Customer, error) {
	customers, err := r.Search(ctx, "email:"+customer.Email)
	if err != nil {
		return nil, fmt.Errorf("error searching for customer: %w", err)
	}

	shopifyCustomer := new(Customer)
	if len(customers) == 0 {
		shopifyCustomer, err = r.Create(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error creating customer: %w", err)
		}
	} else {
		customer.Id = customers[0].Id
		shopifyCustomer, err = r.Update(ctx, customer)
		if err != nil {
			return nil, fmt.Errorf("error updating customer: %w", err)
		}
	}

	return shopifyCustomer, nil
}
