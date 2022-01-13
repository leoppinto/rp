package model

type Product struct {
	ID   string
	Name string
}

func NewProduct() *Product {
	product := Product{}
	product.ID = "AHHSA"
	return &Product{}
}

type Products struct {
	Product []Product
}

func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, *product)
}

func NewProducts() *Products {
	return &Products{}
}
