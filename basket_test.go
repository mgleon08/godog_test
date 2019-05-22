package main

import "github.com/DATA-DOG/godog"

type shopping struct {
	shelf *Shelf
}

func (sh *shopping) addProduct(productName string, price float64) (err error) {
	sh.shelf.AddProduct(productName, price)
	return
}

func iAddTheToTheBasket(arg1 string) error {
	return godog.ErrPending
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
	})

	s.Step(`^there is a "([^"]*)", which costs £(\d+)$`, sh.addProduct)
	s.Step(`^I add the "([^"]*)" to the basket$`, iAddTheToTheBasket)
	s.Step(`^I should have (\d+) products in the basket$`, iShouldHaveProductsInTheBasket)
	s.Step(`^the overall basket price should be £(\d+)$`, theOverallBasketPriceShouldBe)
}
