package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

func NewCheckout(myCheckout checkout.ICheckout) IHandler {
	return &checkoutHandler{
		myCheckout: myCheckout,
	}
}

type checkoutHandler struct {
	myCheckout checkout.ICheckout
}

// Setup implements IHandler.
func (ch *checkoutHandler) Setup(r *gin.Engine) {
	r.GET("/checkout", ch.Get)
	r.POST("/checkout/scan/:sku", ch.Scan)
}

var ErrSKUNotFound error = errors.New("sku not found")

// Scan
// @Summary Scan an item
// @Description Scans an item by it's provided SKU
// @Tags checkout
// @Produce plain
// @Param sku path string true "SKU"
// @Success 200 {string} OK
// @Failure 404 {object} ErrorRes
// @Router /checkout/scan/{sku} [post]
func (ch *checkoutHandler) Scan(c *gin.Context) {
	sku := c.Params.ByName("sku")

	if err := ch.myCheckout.Scan(sku); err != nil {
		c.JSON(http.StatusNotFound, NewErrorRes(ErrSKUNotFound))
		return
	}

	c.String(http.StatusOK, "OK")
}

// Get Total
// @Summary Get the total of the checkout
// @Description Returns the total value of the checkout including discounts
// @Tags checkout
// @Produce plain
// @Success 200 {integer} 0
// @Failure 404 {object} ErrorRes
// @Router /checkout [get]
func (ch *checkoutHandler) Get(c *gin.Context) {
	val, err := ch.myCheckout.GetTotal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorRes(ErrUnexpected))
		return
	}

	c.String(http.StatusOK, strconv.Itoa(val))
}
