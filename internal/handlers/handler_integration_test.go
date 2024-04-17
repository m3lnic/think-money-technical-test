package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
)

func CreateBaseService(handlers ...handlers.IHandler) (
	*gin.Engine,
	checkout.ICatalogueRepository,
	checkout.IDiscountCatalogueRepository,
	checkout.ICheckout,
) {
	myCatalogue := checkout.NewCatalogue()
	myDiscountCatalogue := checkout.NewDiscountCatalogue(myCatalogue)

	myCatalogue.Create("A", checkout.NewItem("Pineapples", 50))
	myCatalogue.Create("B", checkout.NewItem("Waffles", 30))
	myCatalogue.Create("C", checkout.NewItem("Bacon", 20))
	myCatalogue.Create("D", checkout.NewItem("Maple Syrup", 15))

	myDiscountCatalogue.Create("A", checkout.NewDiscount(3, 130))
	myDiscountCatalogue.Create("B", checkout.NewDiscount(2, 45))

	myCheckout := checkout.New(myCatalogue, myDiscountCatalogue)

	myEngine := gin.New()

	return myEngine, myCatalogue, myDiscountCatalogue, myCheckout
}
