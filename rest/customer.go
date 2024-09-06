package rest

import (
	"context"

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
