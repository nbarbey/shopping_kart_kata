package pkg

type ShoppingCart struct {
}

func (c ShoppingCart) String() string {
	return `
--------------------------------------------
| Product Name | Price with VAT | Quantity |
|--------------|----------------|----------|
|------------------------------------------|
| Promotion:                               |
|------------------------------------------|
| Total products: 0                        |
| Total price: 0.00 €                      |
--------------------------------------------`
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}
