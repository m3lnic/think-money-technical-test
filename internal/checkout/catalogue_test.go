package checkout_test

import (
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

// > Step by step test
// > Parent test is parallisable, children tests aren't
func TestNewCatalogue(t *testing.T) {
	t.Parallel()

	itemTestName := "Test"

	myCatalogue := checkout.NewCatalogue()

	t.Run("creates new item", func(t *testing.T) {
		if _, err := myCatalogue.Create("A", checkout.NewItem(itemTestName, 50)); err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}
		if _, err := myCatalogue.Create("A", checkout.NewItem(itemTestName, 50)); err == nil {
			t.Errorf("expected err(%+v), got nil", err)
		}
	})

	t.Run("reads new item", func(t *testing.T) {
		fetchedVal, err := myCatalogue.Read("A")
		if err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}
		if fetchedVal.GetName() != itemTestName {
			t.Errorf("expected string(%s), got string(%s)", itemTestName, fetchedVal.GetName())
		}
	})
}
