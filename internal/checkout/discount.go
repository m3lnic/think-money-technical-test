package checkout

import "encoding/json"

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
	return 0, quantity % i.Quantity
}

func (i Discount) ToString() string {
	myString, _ := json.Marshal(i)
	return string(myString)
}
