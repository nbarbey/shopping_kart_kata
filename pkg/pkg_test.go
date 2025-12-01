package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCart(t *testing.T) {
	cart := NewShoppingCart()

	result := cart.String()

	expected := `
--------------------------------------------
| Product Name | Price with VAT | Quantity |
|--------------|----------------|----------|
|------------------------------------------|
| Promotion:                               |
|------------------------------------------|
| Total products: 0                        |
| Total price: 0.00 €                      |
--------------------------------------------`
	assert.Equal(t, expected, result)
}
