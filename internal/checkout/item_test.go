package checkout_test

import (
	"fmt"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
)

type ItemToStringTestConfig struct {
	Item           checkout.IItem
	ExpectedString string
}

func TestItemToString(t *testing.T) {
	t.Parallel()

	tests := []ItemToStringTestConfig{
		{
			Item:           checkout.NewItem("", 0),
			ExpectedString: "{\"name\":\"\",\"unitPrice\":0}",
		},
		{
			Item:           checkout.NewItem("Pineapples", 50),
			ExpectedString: "{\"name\":\"Pineapples\",\"unitPrice\":50}",
		},
	}

	for key, current := range tests {
		t.Run(fmt.Sprintf("TestItemToString[%d]", key), func(t *testing.T) {
			expectedItemString := current.ExpectedString
			itemString := current.Item.ToString()

			if itemString != expectedItemString {
				t.Errorf("expected string(%s), got string(%s)", expectedItemString, itemString)
			}
		})
	}
}

func TestItemGetUnitPrice(t *testing.T) {
	t.Parallel()

	unitPrice := 50
	item := checkout.NewItem("Pineapples", unitPrice)
	priceOut := item.GetUnitPrice()

	if priceOut != unitPrice {
		t.Errorf("expected unitPrice(%d), got unitPrice(%d)", unitPrice, priceOut)
	}
}

func TestItemGetName(t *testing.T) {
	t.Parallel()

	itemName := "Pineapples"
	item := checkout.NewItem(itemName, 0)
	nameOut := item.GetName()

	if nameOut != itemName {
		t.Errorf("expected name(%s), got name(%s)", itemName, nameOut)
	}
}
