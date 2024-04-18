package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/stretchr/testify/assert"
)

func TestCatalogueSentenceParser(t *testing.T) {
	invalidSentence := "This is an invalid sentence"
	_, _, _, err := checkout.ParseCatalogueItemOrDiscountFromSentence(invalidSentence)
	assert.ErrorIs(t, err, checkout.ErrInvalidInput)

	validSentenceOne := "Pancakes cost 60"
	name, cost, quantity, err := checkout.ParseCatalogueItemOrDiscountFromSentence(validSentenceOne)
	assert.Nil(t, err)
	assert.Equal(t, "Pancakes", name)
	assert.Equal(t, 60, cost)
	assert.Equal(t, 0, quantity)

	validSentenceTwo := "90 Waffles cost 60"
	name, cost, quantity, err = checkout.ParseCatalogueItemOrDiscountFromSentence(validSentenceTwo)
	assert.Nil(t, err)
	assert.Equal(t, "Waffles", name)
	assert.Equal(t, 60, cost)
	assert.Equal(t, 90, quantity)

	validSentenceThree := "90 Loaves of Bread cost 60"
	name, cost, quantity, err = checkout.ParseCatalogueItemOrDiscountFromSentence(validSentenceThree)
	assert.Nil(t, err)
	assert.Equal(t, "Loaves of Bread", name)
	assert.Equal(t, 60, cost)
	assert.Equal(t, 90, quantity)
}
