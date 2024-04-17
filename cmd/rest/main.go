package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
)

const (
	DEFAULT_IP   string = "0.0.0.0"
	DEFAULT_PORT int    = 4000
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

	myCheckout := checkout.New(myCatalogue, myDiscountCatalogue)

	r := gin.New()

	handlers.NewCheckout(myCheckout).Setup(r)

	if err := r.Run(fmt.Sprintf("%s:%d", DEFAULT_IP, DEFAULT_PORT)); err != nil {
		panic(err)
	}
}
