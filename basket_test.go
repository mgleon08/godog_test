package main

import "github.com/DATA-DOG/godog"

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

func iShouldHaveProductsInTheBasket(arg1 int) error {
	return godog.ErrPending
}

func theOverallBasketPriceShouldBe(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	sh := &shopping{}
	s.BeforeScenario(func(interface{}) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()
	})

	s.Step(`^there is a "([^"]*)", which costs £(\d+)$`, sh.addProduct)
	s.Step(`^I add the "([^"]*)" to the basket$`, sh.addToBasket)
	s.Step(`^I should have (\d+) products in the basket$`, iShouldHaveProductsInTheBasket)
	s.Step(`^the overall basket price should be £(\d+)$`, theOverallBasketPriceShouldBe)
}
