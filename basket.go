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

// GetBasketSize returns the number of products in the basket
func (b *Basket) GetBasketSize() int {
	return len(b.products)
}

// GetBasketTotal returns the total price of the items in the basket
// inclusive of VAT of 20% and delivery costs
func (b *Basket) GetBasketTotal() float64 {
	basketTotal := 0.00
	shippingPrice := 0.00

	for _, value := range b.products {
		basketTotal += value
	}
	basketTotal = basketTotal * 1.2

	if basketTotal <= 10 {
		shippingPrice = 3
	} else {
		shippingPrice = 2
	}

	return basketTotal + shippingPrice
}
