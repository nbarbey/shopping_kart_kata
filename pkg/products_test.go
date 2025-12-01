package pkg

import (
	"testing"

	money "github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductsEmpty(t *testing.T) {
	products := NewProducts()

	assert.Zero(t, products.Len())
}

func TestProducts_with_one_product(t *testing.T) {
	products := NewProducts()

	products.Add(Product{Name: "Iceberg"})

	assert.Equal(t, 1, products.Len())
	assert.Equal(t, Product{Name: "Iceberg"}, products.Get("Iceberg"))
}

func TestProducts_final_price(t *testing.T) {
	products := NewProducts()

	products.Add(Product{
		Name:              "Iceberg",
		Cost:              money.New(155, money.EUR),
		PercentageRevenue: 15,
		TaxPecentage:      21,
	})

	equals, err := products.
		Get("Iceberg").
		FinalPrice().
		Equals(money.New(217, money.EUR))
	require.NoError(t, err)
	assert.True(t, equals)
}
