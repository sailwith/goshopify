package rest

import (
	"context"
	"math/rand"
	"testing"

	"github.com/sailwith/goshopify/test"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestCustomerCreate(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	customer, err := rest.CustomerCreate(context.Background(), Customer{
		Email: randString(8) + "@gmail.com",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(customer.Id, customer.Email)
}

func TestCustomerSearch(t *testing.T) {
	client := test.NewClient()
	rest := NewREST(client)
	customers, err := rest.CustomerSearch(context.Background(), "email:111@gmail.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(customers))
}

func TestCustomerUpdate(t *testing.T) {
	client := test.NewClient()
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
	client := test.NewClient()
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
