package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

func TestNewCatalogue(t *testing.T) {
	itemTestName := "Test"

	myCatalogue := checkout.NewCatalogue()
	newItem := checkout.NewItem(itemTestName, 50)

	_, err := myCatalogue.Create("A", newItem)
	if err != nil {
		t.Errorf("expected nil, got error(%+v)", err)
	}

	fetchedVal, err := myCatalogue.Read("A")
	if err != nil {
		t.Errorf("expected nil, got error(%+v)", err)
	}
	if fetchedVal.GetName() != itemTestName {
		t.Errorf("expected string(%s), got string(%s)", itemTestName, fetchedVal.GetName())
	}
}
