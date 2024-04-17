package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/repository"
	"github.com/stretchr/testify/assert"
)

// > Step by step test
func TestCheckoutScan(t *testing.T) {
	t.Parallel()

	// > We could use mocking here, but as catalogue is a memory repo
	// > I think it should be ok to use this as is for now
	myCat := checkout.NewCatalogue()
	pineapple, _ := myCat.Create("A", checkout.NewItem("Pineapples", 50))
	// waffle, _ := myCat.Create("B", checkout.NewItem("Waffles", 30))
	myCat.Create("B", checkout.NewItem("Waffles", 30))

	myDiscountCatalogue := checkout.NewDiscountCatalogue(myCat)
	myDiscountCatalogue.Create("A", checkout.NewDiscount(3, 130))

	myCheckout := checkout.New(myCat)

	t.Run("validate total is initally 0", func(t *testing.T) {
		val, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		assert.Equal(t, val, 0)
	})

	t.Run("when sku scanned, total is calculated correctly", func(t *testing.T) {
		err := myCheckout.Scan("A")
		assert.Nil(t, err)

		val, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		assert.Equal(t, val, pineapple.GetUnitPrice())

		err = myCheckout.Scan("A")
		assert.Nil(t, err)

		newVal, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		assert.Equal(t, newVal, pineapple.GetUnitPrice()*2)
	})

	t.Run("errors when sku not found", func(t *testing.T) {
		err := myCheckout.Scan("B")
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repository.ErrKeyNotFound)
	})

	t.Run("when sku has been removed before getting the total", func(t *testing.T) {
		myCat.Create("B", checkout.NewItem("Waffles", 25))
		myCheckout.Scan("B")
		myCat.Delete("B")

		_, err := myCheckout.GetTotal()
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repository.ErrKeyNotFound)
	})
}
