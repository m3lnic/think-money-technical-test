package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/repository"
	"github.com/stretchr/testify/assert"
)

// > Step by step test
// > Parent test is parallisable, children tests aren't
func TestNewCatalogue(t *testing.T) {
	t.Parallel()

	itemTestName := "Test"

	myCatalogue := checkout.NewCatalogue()

	t.Run("creates new item", func(t *testing.T) {
		_, err := myCatalogue.Create("A", checkout.NewItem(itemTestName, 50))
		assert.Nil(t, err)

		_, secErr := myCatalogue.Create("A", checkout.NewItem(itemTestName, 50))
		assert.NotNil(t, secErr)
		assert.ErrorIs(t, secErr, repository.ErrKeyAlreadyExists)
	})

	t.Run("reads new item", func(t *testing.T) {
		fetchedVal, err := myCatalogue.Read("A")
		assert.Nil(t, err)
		assert.Equal(t, fetchedVal.GetName(), itemTestName)
	})
}

func TestGetByItemName(t *testing.T) {
	itemName := "Test"
	myCatalogue := checkout.NewCatalogue()
	myCatalogue.Create("A", checkout.NewItem(itemName, 0))

	_, _, err := myCatalogue.GetByItemName("NotTest")
	assert.ErrorIs(t, err, repository.ErrKeyNotFound)

	sku, item, err := myCatalogue.GetByItemName(itemName)
	assert.Nil(t, err)
	assert.Equal(t, "A", sku)
	assert.Equal(t, itemName, item.GetName())
}
