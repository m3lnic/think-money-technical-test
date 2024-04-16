package main

import "github.com/m3lnic/think-money-technical-test/internal/checkout"

func main() {
	myCatalogue := checkout.NewCatalogue()
	myCatalogue.Create("A", checkout.NewItem("Pineapples", 50))
	myCatalogue.Create("B", checkout.NewItem("Waffles", 25))
	myCatalogue.Create("C", checkout.NewItem("Bacon", 100))

	myCheckout := checkout.New(myCatalogue)
	myCheckout.Scan("A")
	myCheckout.Scan("B")
	myCheckout.Scan("A")
}
