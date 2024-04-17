package checkout_test

import (
	"fmt"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/stretchr/testify/assert"
)

type ItemToStringTestConfig struct {
	Item           checkout.Item
	ExpectedString string
}

func TestItemToString(t *testing.T) {
	t.Parallel()

	tests := []ItemToStringTestConfig{
		{
			Item:           *checkout.NewItem("", 0),
			ExpectedString: "{\"name\":\"\",\"unitPrice\":0}",
		},
		{
			Item:           *checkout.NewItem("Pineapples", 50),
			ExpectedString: "{\"name\":\"Pineapples\",\"unitPrice\":50}",
		},
	}

	for key, current := range tests {
		t.Run(fmt.Sprintf("TestItemToString[%d]", key), func(t *testing.T) {
			expectedItemString := current.ExpectedString
			itemString := current.Item.ToString()

			assert.Equal(t, itemString, expectedItemString)
		})
	}
}

func TestItemGetUnitPrice(t *testing.T) {
	t.Parallel()

	unitPrice := 50
	item := checkout.NewItem("Pineapples", unitPrice)
	assert.Equal(t, item.GetUnitPrice(), unitPrice)
}

func TestItemGetName(t *testing.T) {
	t.Parallel()

	itemName := "Pineapples"
	item := checkout.NewItem(itemName, 0)
	assert.Equal(t, item.GetName(), itemName)
}
