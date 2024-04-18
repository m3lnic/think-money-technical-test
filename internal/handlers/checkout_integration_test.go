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

func initializeCheckoutTest() (
	*gin.Engine,
	checkout.ICatalogueRepository,
	checkout.IDiscountCatalogueRepository,
	checkout.ICheckout,
) {
	engine, cat, discoCat, checkout := CreateBaseService()

	handlers.NewCheckout(checkout).Setup(engine)

	return engine, cat, discoCat, checkout
}

func TestCheckoutHandlerScan(t *testing.T) {
	t.Parallel()

	myEngine, _, _, _ := initializeCheckoutTest()

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

func TestCheckoutHandlerGet(t *testing.T) {
	t.Parallel()

	t.Run("returns correct data", func(t *testing.T) {
		t.Parallel()

		myEngine, _, _, _ := initializeCheckoutTest()

		t.Parallel()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/checkout", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "0", w.Body.String())
	})

	t.Run("when item removed", func(t *testing.T) {
		t.Parallel()

		myEngine, myCatalogue, _, myCheckout := initializeCheckoutTest()

		t.Parallel()

		myCheckout.Scan("A")

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/checkout", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "50", w.Body.String())

		myCatalogue.Delete("A")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/checkout", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, handlers.NewErrorRes(handlers.ErrUnexpected).ToString(), w.Body.String())
	})
}
