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

// Create or Update by SKU
// @Summary Creates or Updates a discount by it's provided SKU
// @Description Creates or updates a discount by it's SKU
// @Tags discount
// @Produce json
// @Param sku path string true "SKU"
// @Param body body CreateOrUpdateDiscountReq true "The data for the discount you'd like to apply"
// @Success 200 {boolean} True
// @Failure 404 {object} ErrorRes
// @Router /discount/{sku} [post]
func (dh *discountHandler) CreateOrUpdateBySKU(c *gin.Context) {
	sku := c.Params.ByName("sku")

	var body CreateOrUpdateDiscountReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorRes(ErrInvalidBody))
		return
	}

	updateOrCreateVal := checkout.NewDiscount(body.Quantity, body.Price)
	if _, err := dh.discountCatalogue.Read(sku); err == nil {
		dh.discountCatalogue.Update(sku, updateOrCreateVal)
	} else {
		if _, err := dh.discountCatalogue.Create(sku, updateOrCreateVal); err != nil {
			c.JSON(http.StatusNotFound, NewErrorRes(ErrSKUNotFound))
			return
		}
	}

	c.JSON(http.StatusOK, true) // > Will convert true to string
}
