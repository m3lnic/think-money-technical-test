package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
	_ "github.com/m3lnic/think-money-technical-test/pkg/docs"
)

const (
	INITIAL_PARSED_SENTENCE string = ""
	DEFAULT_IP              string = "0.0.0.0"
	DEFAULT_PORT            int    = 4000
)

// @title ThinkMoney technical test
// @version 1.0
// @description This is an example repository for the technical test of think money

// @contact.name Melody Nicholls
// @contact.email melody@technode.uk

// @host localhost:4000
// @BasePath /
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
	myCatalogueSentenceParser := checkout.NewCatalogueSentenceParser(myCatalogue, myDiscountCatalogue)

	if INITIAL_PARSED_SENTENCE != "" {
		if err := myCatalogueSentenceParser.Parse(INITIAL_PARSED_SENTENCE); err != nil {
			panic(err)
		}
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handlers.NewCheckout(myCheckout).Setup(r)
	handlers.NewDiscount(myDiscountCatalogue).Setup(r)
	handlers.NewCatalogue(myCatalogueSentenceParser).Setup(r)

	if err := r.Run(fmt.Sprintf("%s:%d", DEFAULT_IP, DEFAULT_PORT)); err != nil {
		panic(err)
	}
}
