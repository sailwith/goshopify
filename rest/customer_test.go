package rest

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/test"
)

func TestCustomerCreate(t *testing.T) {
	client, err := test.NewClient()
	if err != nil {
		t.Error(err)
	}
	rest := NewREST(client)
	customer, err := rest.CustomerCreate(context.Background(), Customer{
		Email: test.RandString(8) + "@gmail.com",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(customer.Id, customer.Email)
}

func TestCustomerSearch(t *testing.T) {
	client, err := test.NewClient()
	if err != nil {
		t.Error(err)
	}
	rest := NewREST(client)
	customers, err := rest.CustomerSearch(context.Background(), "email:111@gmail.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(customers))
}

func TestCustomerUpdate(t *testing.T) {
	client, err := test.NewClient()
	if err != nil {
		t.Error(err)
	}
	rest := NewREST(client)
	c, err := rest.CustomerUpdate(context.Background(), Customer{
		Id:        7233840578864,
		FirstName: "first name",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(c.FirstName)
}

func TestCustomerCreateOrUpdate(t *testing.T) {
	client, err := test.NewClient()
	if err != nil {
		t.Error(err)
	}
	rest := NewREST(client)
	c, err := rest.CustomerCreateOrUpdate(context.Background(), Customer{
		Email:    "111@gmail.com",
		LastName: "last name",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(c.LastName)
}
