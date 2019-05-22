package main

// NewShelf instantiates a new Shelf object
func NewShelf() *Shelf {
	return &Shelf{
		products: make(map[string]float64),
	}
}

// Shelf stores a list of products which are available for purchase type Shelf
type Shelf struct {
	products map[string]float64
}

// AddProduct adds a product to the shelf
func (s *Shelf) AddProduct(productName string, price float64) {
	s.products[productName] = price
}
