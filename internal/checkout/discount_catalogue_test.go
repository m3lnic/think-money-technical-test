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

	myCatalogue := checkout.NewDiscountCatalogue()

	t.Run("creates new item", func(t *testing.T) {
		_, err := myCatalogue.Create("A", checkout.NewDiscount(1, 2))
		assert.Nil(t, err)

		_, secErr := myCatalogue.Create("A", checkout.NewDiscount(1, 2))
		assert.NotNil(t, secErr)
		assert.ErrorIs(t, secErr, repository.ErrKeyAlreadyExists)
	})

	t.Run("reads new item", func(t *testing.T) {
		fetchedVal, err := myCatalogue.Read("A")
		assert.Nil(t, err)

		expectedTotal, remaining := fetchedVal.QualifiesFor(1)
		assert.Equal(t, expectedTotal, 2)
		assert.Equal(t, remaining, 0)
	})
}
