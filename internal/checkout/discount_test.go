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

	discount := checkout.NewDiscount(4, 50)

	discountTotal, remainingItems := discount.QualifiesFor(0)
	assert.Equal(t, discountTotal, 0)
	assert.Equal(t, remainingItems, 0)

	discountTotal, remainingItems = discount.QualifiesFor(2)
	assert.Equal(t, discountTotal, 0)
	assert.Equal(t, remainingItems, 2)
}
