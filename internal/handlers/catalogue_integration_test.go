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

func initializeCatalogueTest() (
	*gin.Engine,
	checkout.ICatalogueRepository,
	checkout.IDiscountCatalogueRepository,
	checkout.ICheckout,
) {
	engine, cat, discoCat, myCheckout := CreateBaseService()

	sentenceParser := checkout.NewCatalogueSentenceParser(cat, discoCat)
	handlers.NewCatalogue(sentenceParser).Setup(engine)

	return engine, cat, discoCat, myCheckout
}

func TestParseBySentence(t *testing.T) {
	t.Parallel()

	myEngine, _, _, _ := initializeCatalogueTest()

	t.Run("when valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(handlers.ParseBySentenceReq{
			Sentence: "Pineapples cost 50. 2 Pineapples cost 75.",
		})

		req, _ := http.NewRequest("POST", "/catalogue/by-sentence", bytes.NewBuffer(body))
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("when sentence invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(handlers.ParseBySentenceReq{
			Sentence: "Pineapples are 50. 2 Pineapples cost 75.",
		})

		req, _ := http.NewRequest("POST", "/catalogue/by-sentence", bytes.NewBuffer(body))
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, handlers.NewErrorRes(handlers.ErrInvalidBody).ToString(), w.Body.String())
	})

	t.Run("when item doesn't exist in the catalogue", func(t *testing.T) {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(handlers.ParseBySentenceReq{
			Sentence: "Chicken cost 50. 2 Pineapples cost 75.",
		})

		req, _ := http.NewRequest("POST", "/catalogue/by-sentence", bytes.NewBuffer(body))
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, handlers.NewErrorRes(handlers.ErrSKUNotFound).ToString(), w.Body.String())
	})

	t.Run("when invalid body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/catalogue/by-sentence", nil)
		myEngine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, handlers.NewErrorRes(handlers.ErrInvalidBody).ToString(), w.Body.String())
	})
}
