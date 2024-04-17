package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/repository"
	"github.com/stretchr/testify/assert"
)

// > Step by step test
// > Parent test is parallisable, children tests aren't
func TestNewDiscountCatalogue(t *testing.T) {
	t.Parallel()

	myCatalogue := checkout.NewCatalogue()
	myCatalogue.Create("A", checkout.NewItem("Pineapples", 20))

	myDiscountCatalogue := checkout.NewDiscountCatalogue(myCatalogue)

	t.Run("creates new item", func(t *testing.T) {
		data, err := myDiscountCatalogue.Create("A", checkout.NewDiscount(1, 2))
		assert.Nil(t, err)
		assert.Equal(t, data.Price, 2)

		_, secErr := myDiscountCatalogue.Create("A", checkout.NewDiscount(1, 2))
		assert.NotNil(t, secErr)
		assert.ErrorIs(t, secErr, repository.ErrKeyAlreadyExists)
	})

	t.Run("reads new item", func(t *testing.T) {
		fetchedVal, err := myDiscountCatalogue.Read("A")
		assert.Nil(t, err)

		expectedTotal, remaining := fetchedVal.QualifiesFor(1)
		assert.Equal(t, expectedTotal, 2)
		assert.Equal(t, remaining, 0)
	})

	t.Run("when new item doesn't exist in the catalogue", func(t *testing.T) {
		data, err := myDiscountCatalogue.Create("B", checkout.NewDiscount(1, 2))
		assert.Nil(t, data)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repository.ErrKeyNotFound)
	})
}
