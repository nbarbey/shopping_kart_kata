package pkg

import (
	"fmt"

	money "github.com/Rhymond/go-money"
)

type Item string

type Quantity int

type Items map[Item]Quantity

type ShoppingCart struct {
	products Products
	items    Items
}

func (c ShoppingCart) String() string {
	const (
		header = `--------------------------------------------
| Product Name | Price with VAT | Quantity |
|--------------|----------------|----------|`

		promotion = `|------------------------------------------|
| Promotion:                               |
|------------------------------------------|`
	)

	return fmt.Sprintf(`
%s
%s
| Total products: 0                        |
| Total price: 0.00 €                      |
--------------------------------------------`, header, promotion)
}

func NewShoppingCart() *ShoppingCart {
	cart := &ShoppingCart{
		products: NewProducts(),
		items:    make(Items),
	}
	cart.products.Add(Product{
		Name:              "Iceberg",
		Cost:              money.New(155, money.EUR),
		PercentageRevenue: 15,
		TaxPecentage:      21,
	})

	return cart
}

func (c *ShoppingCart) Add(s string) {
	c.items[Item(s)]++
}
