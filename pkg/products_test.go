package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
