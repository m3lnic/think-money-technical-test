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
	waffle, _ := myCat.Create("B", checkout.NewItem("Waffles", 30))

	myDiscountCatalogue := checkout.NewDiscountCatalogue(myCat)
	pineappleDiscount, _ := myDiscountCatalogue.Create("A", checkout.NewDiscount(3, 130))

	myCheckout := checkout.New(myCat)

	t.Run("validate total is initally 0", func(t *testing.T) {
		val, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		assert.Equal(t, val, 0)
	})

	t.Run("when sku scanned, total is calculated correctly", func(t *testing.T) {
		err := myCheckout.Scan("A")
		assert.Nil(t, err)

		runningTotal := 0

		val, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		runningTotal += pineapple.GetUnitPrice()
		assert.Equal(t, val, runningTotal)

		err = myCheckout.Scan("A")
		assert.Nil(t, err)

		newVal, err := myCheckout.GetTotal()
		assert.Nil(t, err)
		runningTotal += pineapple.GetUnitPrice()
		assert.Equal(t, newVal, runningTotal)

		err = myCheckout.Scan("B")
		assert.Nil(t, err)

		newVal, err = myCheckout.GetTotal()
		assert.Nil(t, err)
		runningTotal += waffle.GetUnitPrice()
		assert.Equal(t, newVal, runningTotal)

		err = myCheckout.Scan("A")
		assert.Nil(t, err)

		newVal, err = myCheckout.GetTotal()
		assert.Nil(t, err)
		myNewTotal := pineappleDiscount.Price + waffle.GetUnitPrice()
		assert.Equal(t, newVal, myNewTotal)
	})

	t.Run("errors when sku not found", func(t *testing.T) {
		err := myCheckout.Scan("C")
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repository.ErrKeyNotFound)
	})

	t.Run("when sku has been removed before getting the total", func(t *testing.T) {
		myCat.Create("C", checkout.NewItem("Cheesy Puffs", 25))
		myCheckout.Scan("C")
		myCat.Delete("C")

		_, err := myCheckout.GetTotal()
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repository.ErrKeyNotFound)
	})
}
