package main

import (
	"log"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

func main() {
	myCatalogue := checkout.NewCatalogue()
	myCatalogue.Create("A", checkout.NewItem("Pineapples", 50))
	myCatalogue.Create("B", checkout.NewItem("Waffles", 30))
	myCatalogue.Create("C", checkout.NewItem("Bacon", 20))
	myCatalogue.Create("D", checkout.NewItem("Maple Syrup", 15))

	myDiscountCatalogue := checkout.NewDiscountCatalogue(myCatalogue)
	myDiscountCatalogue.Create("A", checkout.NewDiscount(3, 130))
	myDiscountCatalogue.Create("B", checkout.NewDiscount(2, 45))

	myCheckout := checkout.New(myCatalogue)
	myCheckout.Scan("A")
	myCheckout.Scan("B")
	myCheckout.Scan("A")
	myCheckout.Scan("A")
	myCheckout.Scan("C")
	myCheckout.Scan("D")
	myCheckout.Scan("B")

	total, err := myCheckout.GetTotal()
	if err != nil {
		panic(err)
	}

	log.Printf("Total: %d", total)
}
