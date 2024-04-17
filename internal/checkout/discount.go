package checkout

import (
	"encoding/json"
	"math"
)

// > Mainly used for reference
type IDiscount interface {
	// > Out - discountTotal, remainingItems
	QualifiesFor(quantity int) (int, int)
	ToString() string
}

// > Function means that we must always pass these values when creating an item
func NewDiscount(quantity, price int) *Discount {
	return &Discount{
		Quantity: quantity,
		Price:    price,
	}
}

type Discount struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

func (i Discount) QualifiesFor(quantity int) (int, int) {
	totalCost := int(math.Floor(float64(quantity) / float64(i.Quantity)))
	remainingQuantity := quantity % i.Quantity

	return totalCost * i.Price, remainingQuantity
}

func (i Discount) ToString() string {
	myString, _ := json.Marshal(i)
	return string(myString)
}
