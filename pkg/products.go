package pkg

type Product struct {
	Name string
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

func NewProducts() Products {
	return make(Products, 0)
}
