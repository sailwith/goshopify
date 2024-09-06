package helper

import (
	"context"
	"testing"

	"github.com/sailwith/goshopify/rest"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateCustomer(t *testing.T) {
	h := newHelper(t)
	c, err := h.CreateOrUpdateCustomer(context.Background(), rest.Customer{
		Email:    "111@gmail.com",
		LastName: "last name",
	})
	if assert.NoError(t, err) {
		t.Log(c.LastName)
	}
}
