package model

type Product struct {
	ID   string
	Name string
}

type Products struct {
	Product []Product
}

func (p *Products) Add(product Product) {
	p.Product = append(p.Product, product)
}