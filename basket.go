package main

// NewBasket instantiates a new Basket object
func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]float64),
	}
}

// Basket stores a list of products which a user wishes to purchase
type Basket struct {
	products map[string]float64
}

// GetBasketPrice calculates the total basket price without VAT or shipping
func (b *Basket) AddItem(productName string, price float64) {
	b.products[productName] = price
}
