package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

func NewDiscount(discountCatalog checkout.IDiscountCatalogueRepository) IHandler {
	return &discountHandler{
		discountCatalogue: discountCatalog,
	}
}

type discountHandler struct {
	discountCatalogue checkout.IDiscountCatalogueRepository
}

// Setup implements IHandler.
func (dh *discountHandler) Setup(r *gin.Engine) {
	r.POST("/discount/:sku", dh.CreateOrUpdateBySKU)
}

type CreateOrUpdateDiscountReq struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

func (dh *discountHandler) CreateOrUpdateBySKU(c *gin.Context) {
	// sku := c.Params.ByName("sku")

	var body CreateOrUpdateDiscountReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorRes(ErrInvalidBody))
		return
	}

	c.JSON(http.StatusOK, true) // > Will convert true to string
}
