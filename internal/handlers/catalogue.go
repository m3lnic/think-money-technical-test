package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/repository"
)

func NewCatalogue(sentenceParser checkout.ICatalogueSentenceParser) IHandler {
	return &catalogueHandler{
		sentenceParser: sentenceParser,
	}
}

type catalogueHandler struct {
	sentenceParser checkout.ICatalogueSentenceParser
}

// Setup implements IHandler.
func (ch *catalogueHandler) Setup(r *gin.Engine) {
	r.POST("/catalogue/by-sentence", ch.ParseBySentence)
}

type ParseBySentenceReq struct {
	Sentence string `json:"sentence"`
}

func (ch *catalogueHandler) ParseBySentence(c *gin.Context) {
	var body ParseBySentenceReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorRes(ErrInvalidBody))
		return
	}

	if err := ch.sentenceParser.Parse(body.Sentence); err != nil {
		if errors.Is(err, repository.ErrKeyNotFound) {
			c.JSON(http.StatusNotFound, NewErrorRes(ErrSKUNotFound))
			return
		}

		if errors.Is(err, checkout.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, NewErrorRes(ErrInvalidBody))
			return
		}
	}

	c.String(http.StatusOK, "OK")
}
