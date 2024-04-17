package checkout_test

import (
	"errors"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/m3lnic/think-money-technical-test/internal/repository"
)

// > Step by step test
func TestCheckoutScan(t *testing.T) {
	t.Parallel()

	// > Should really use mocking here
	myCat := checkout.NewCatalogue()
	pineapple, _ := myCat.Create("A", checkout.NewItem("Pineapples", 50))
	// waffle, _ := myCat.Create("B", checkout.NewItem("Waffles", 25))
	// bacon, _ := myCat.Create("C", checkout.NewItem("Bacon", 100))

	myCheckout := checkout.New(myCat)

	t.Run("validate total is initally 0", func(t *testing.T) {
		if val, _ := myCheckout.GetTotal(); val != 0 {
			t.Errorf("expected int(0), got int(%d)", val)
		}
	})

	t.Run("when sku scanned, total is calculated correctly", func(t *testing.T) {
		err := myCheckout.Scan("A")
		if err != nil {
			t.Errorf("expected nil, got err(%+v)", err)
		}

		if val, _ := myCheckout.GetTotal(); val != pineapple.GetUnitPrice() {
			t.Errorf("expected int(%d), got int(%d)", pineapple.GetUnitPrice(), val)
		}

		err = myCheckout.Scan("A")
		if err != nil {
			t.Errorf("expected nil, got err(%+v)", err)
		}

		if val, _ := myCheckout.GetTotal(); val != pineapple.GetUnitPrice()*2 {
			t.Errorf("expected int(%d), got int(%d)", pineapple.GetUnitPrice()*2, val)
		}
	})

	t.Run("errors when sku not found", func(t *testing.T) {
		err := myCheckout.Scan("B")
		if err == nil {
			t.Errorf("expected err(%+v), got nil", repository.ErrKeyNotFound)
		}
		if !errors.Is(err, repository.ErrKeyNotFound) {
			t.Errorf("expected err(%+v), got err(%+v)", repository.ErrKeyNotFound, err)
		}
	})

	t.Run("when sku has been removed before getting the total", func(t *testing.T) {
		myCat.Create("B", checkout.NewItem("Waffles", 25))
		myCheckout.Scan("B")
		myCat.Delete("B")

		if _, err := myCheckout.GetTotal(); err != nil {
			if !errors.Is(err, repository.ErrKeyNotFound) {
				t.Errorf("expected err(%+v), got err(%+v)", repository.ErrKeyNotFound, err)
			}
		} else {
			t.Errorf("expected err(%+v), got nil", repository.ErrKeyNotFound)
		}
	})
}
