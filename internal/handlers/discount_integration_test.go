package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/handlers"
	"github.com/stretchr/testify/assert"
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
	t.Run("when body is invalid", func(t *testing.T) {
		t.Parallel()

		myEngine, _, _, _ := initializeDiscountTest()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/discount/A", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, handlers.NewErrorRes(handlers.ErrInvalidBody).ToString(), w.Body.String())
	})

	t.Run("when body is valid", func(t *testing.T) {
		t.Parallel()

		myEngine, _, _, _ := initializeDiscountTest()

		t.Run("when catalogue item does not exist", func(t *testing.T) {
			myBody := handlers.CreateOrUpdateDiscountReq{
				Quantity: 10,
				Price:    20,
			}

			ioBody, _ := json.Marshal(myBody)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/discount/Z", bytes.NewBuffer(ioBody))
			myEngine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusNotFound, w.Code)
			assert.Equal(t, handlers.NewErrorRes(handlers.ErrSKUNotFound).ToString(), w.Body.String())
		})

		t.Run("when updated discount does not exist", func(t *testing.T) {
			myBody := handlers.CreateOrUpdateDiscountReq{
				Quantity: 10,
				Price:    20,
			}

			ioBody, _ := json.Marshal(myBody)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/discount/A", bytes.NewBuffer(ioBody))
			myEngine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "true", w.Body.String())
		})

		t.Run("when discount exists", func(t *testing.T) {
			myBody := handlers.CreateOrUpdateDiscountReq{
				Quantity: 10,
				Price:    20,
			}

			ioBody, _ := json.Marshal(myBody)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/discount/A", bytes.NewBuffer(ioBody))
			myEngine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "true", w.Body.String())
		})
	})
}
