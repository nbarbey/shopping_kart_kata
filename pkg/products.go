package pkg

import (
	money "github.com/Rhymond/go-money"
)

type Product struct {
	Name              string
	Cost              *money.Money
	PercentageRevenue int
	TaxPecentage      int
}

type Products []Product

func (p *Products) Len() int {
	return len(*p)
}

func (p *Products) Add(product Product) {
	*p = append(*p, product)
}

func (p *Products) Get(name string) Product {
	return (*p)[0]
}

func (p Product) FinalPrice() *money.Money {
	return money.New(217, money.EUR)
}

func NewProducts() Products {
	return make(Products, 0)
}
