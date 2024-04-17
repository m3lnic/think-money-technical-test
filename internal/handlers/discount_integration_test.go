package handlers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
)

func initializeDiscountTest() (
	*gin.Engine,
	checkout.ICatalogueRepository,
	checkout.IDiscountCatalogueRepository,
	checkout.ICheckout,
) {
	engine, cat, discoCat, checkout := CreateBaseService()

	handlers.NewDiscount(discoCat).Setup(engine)

	return engine, cat, discoCat, checkout
}

func TestDiscountHandlerCreateOrUpdate(t *testing.T) {
	t.Run("when sku does not exist", func(t *testing.T) {

	})
}
