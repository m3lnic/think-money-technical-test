package checkout_test

import (
	"fmt"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/checkout"
	"github.com/stretchr/testify/assert"
)

type DiscountToStringTestConfig struct {
	Discount       checkout.Discount
	ExpectedString string
}

func TestDiscountToString(t *testing.T) {
	t.Parallel()

	tests := []DiscountToStringTestConfig{
		{
			Discount:       *checkout.NewDiscount(0, 0),
			ExpectedString: "{\"quantity\":0,\"price\":0}",
		},
		{
			Discount:       *checkout.NewDiscount(2, 30),
			ExpectedString: "{\"quantity\":2,\"price\":30}",
		},
	}

	for key, current := range tests {
		t.Run(fmt.Sprintf("TestItemToString[%d]", key), func(t *testing.T) {
			expectedItemString := current.ExpectedString
			itemString := current.Discount.ToString()

			assert.Equal(t, itemString, expectedItemString)
		})
	}
}

func TestDiscountQualifiesFor(t *testing.T) {
	t.Parallel()

	unitPrice := 50
	item := checkout.NewItem("Pineapples", unitPrice)
	assert.Equal(t, item.GetUnitPrice(), unitPrice)
}
