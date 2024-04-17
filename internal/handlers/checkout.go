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

func (ch *checkoutHandler) Scan(c *gin.Context) {
	sku := c.Params.ByName("sku")

	if err := ch.myCheckout.Scan(sku); err != nil {
		c.JSON(http.StatusNotFound, NewErrorRes(ErrSKUNotFound))
		return
	}

	c.String(http.StatusOK, "OK")
}

func (ch *checkoutHandler) Get(c *gin.Context) {
	val, err := ch.myCheckout.GetTotal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorRes(ErrUnexpected))
		return
	}

	c.String(http.StatusOK, strconv.Itoa(val))
}
