package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
	"github.com/stretchr/testify/assert"
)

func TestCustomerCreate(t *testing.T) {
	rest := newREST(t)
	customer, err := rest.Customer.Create(context.Background(), Customer{
		Email: test.RandString(8) + "@gmail.com",
		EmailMarketingConsent: &EmailMarketingConsent{
			State:      "subscribed",
			OptInLevel: "single_opt_in",
		},
	})
	if assert.NoError(t, err) {
		t.Log(customer.Id, customer.Email)
	}
}

func TestCustomerSearch(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	customers, err := rest.Customer.Search(context.Background(), "email:111@gmail.com")
	if assert.NoError(t, err) {
		t.Log(len(customers))
	}
}

func TestCustomerUpdate(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	c, err := rest.Customer.Update(context.Background(), Customer{
		Id:        7233840578864,
		FirstName: "first name",
	})
	if assert.NoError(t, err) {
		t.Log(c.FirstName)
	}
}

func TestCustomerCreateOrUpdate(t *testing.T) {
	client, err := test.NewClient()
	assert.NoError(t, err)

	rest := NewREST(client)
	c, err := rest.Customer.CreateOrUpdate(context.Background(), Customer{
		Email:    "111@gmail.com",
		LastName: "last name",
	})
	if assert.NoError(t, err) {
		t.Log(c.LastName)
	}
}
