package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

type shopping struct {
	shelf *Shelf
	basket *Basket
}

func (sh *shopping) addProduct(productName string, price float64) (err error) {
	sh.shelf.AddProduct(productName, price)
	return
}

func (sh *shopping) addToBasket(productName string) (err error) {
	sh.basket.AddItem(productName, sh.shelf.GetProductPrice(productName))
	return
}

func (sh *shopping) iShouldHaveProductsInTheBasket(productCount int) error {
	if sh.basket.GetBasketSize() != productCount {
		return fmt.Errorf(
			"expected %d products but there are %d",
			sh.basket.GetBasketSize(),
			productCount,
		)
	}
	return nil
}

func (sh *shopping) theOverallBasketPriceShouldBe(basketTotal float64) error {
	if sh.basket.GetBasketTotal() != basketTotal {
		return fmt.Errorf(
			"expected basket total to be %.2f but it is %.2f",
			basketTotal,
			sh.basket.GetBasketTotal(),
		)
	}
	return nil
}
func FeatureContext(s *godog.Suite) {
	sh := &shopping{}
	s.BeforeScenario(func(interface{}) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()
	})

	s.Step(`^there is a "([^"]*)", which costs £(\d+)$`, sh.addProduct)
	s.Step(`^I add the "([^"]*)" to the basket$`, sh.addToBasket)
	s.Step(`^I should have (\d+) products in the basket$`, sh.iShouldHaveProductsInTheBasket)
	s.Step(`^the overall basket price should be £(\d+)$`, sh.theOverallBasketPriceShouldBe)
}
