package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/stretchr/testify/assert"
)

func TestCatalogueSentenceParserFunc(t *testing.T) {
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

func TestCatalogueSentenceParser(t *testing.T) {
	catalogue := checkout.NewCatalogue()
	catalogue.Create("A", checkout.NewItem("Pineapples", 50))
	catalogue.Create("B", checkout.NewItem("Waffles", 25))
	catalogue.Create("C", checkout.NewItem("Loaf of Bread", 10))

	discountCatalogue := checkout.NewDiscountCatalogue(catalogue)
	discountCatalogue.Create("A", checkout.NewDiscount(3, 100))

	catalogueSentenceParser := checkout.NewCatalogueSentenceParser(catalogue, discountCatalogue)

	inputSentence := "Pineapples cost 80, Pineapples are 5 for 200. Waffles cost 100, Loaf of Bread cost 25. 2 Loaf of Bread cost 40."
	err := catalogueSentenceParser.Parse(inputSentence)
	assert.Nil(t, err)
}
