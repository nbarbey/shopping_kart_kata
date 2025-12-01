package pkg

import (
	"math"

	money "github.com/Rhymond/go-money"
)

type Product struct {
	Name              string
	Cost              *money.Money
	PercentageRevenue Percent
	TaxPecentage      Percent
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

type Percent int

func (p Percent) Float() float64 {
	return float64(p)
}

func (p Percent) Ratio() float64 {
	return p.Float() / 100
}

func (p Percent) Apply(m *money.Money) *money.Money {
	appliedAmount := math.Ceil(floatAmount(m) * p.Ratio())
	add, _ := m.Add(money.New(int64(appliedAmount), m.Currency().Code))
	return add
}

func floatAmount(m *money.Money) float64 {
	return float64(m.Amount())
}

func (p Product) FinalPrice() *money.Money {
	return p.TaxPecentage.Apply(p.PercentageRevenue.Apply(p.Cost))
}

func NewProducts() Products {
	return make(Products, 0)
}
