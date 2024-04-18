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

func initializeCatalogueTest() (
	*gin.Engine,
	checkout.ICatalogueRepository,
	checkout.IDiscountCatalogueRepository,
	checkout.ICheckout,
) {
	engine, cat, discoCat, checkout := CreateBaseService()

	handlers.NewCatalogue().Setup(engine)

	return engine, cat, discoCat, checkout
}

func TestParseBySentence(t *testing.T) {
	myEngine, _, _, _ := initializeCatalogueTest()

	t.Run("/catalogue/by-sentence", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/catalogue/by-sentence", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})
}
