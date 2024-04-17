package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func initializeTest() (
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
	handlers.NewCheckout(myCheckout).Setup(myEngine)

	return myEngine, myCatalogue, myDiscountCatalogue, myCheckout
}

func TestCheckoutHandlerScan(t *testing.T) {
	myEngine, _, _, _ := initializeTest()

	t.Run("/checkout/scan/:sku", func(t *testing.T) {
		t.Run("when sku found", func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/checkout/scan/A", nil)
			myEngine.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
			assert.Equal(t, "OK", w.Body.String())
		})

		t.Run("when sku not found", func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/checkout/scan/Z", nil)
			myEngine.ServeHTTP(w, req)

			assert.Equal(t, 404, w.Code)
			assert.Equal(t, handlers.NewErrorRes(handlers.ErrSKUNotFound).ToString(), w.Body.String())
		})
	})
}
