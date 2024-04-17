package checkout

import "encoding/json"

type IDiscount interface {
	// > Out - totalCost, remainingItems
	QualifiesForDiscount(quantity int) (int, int)
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
	IDiscount `json:"-"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
}

func (i Discount) QualifiesFor(quantity int) (int, int) {
	return 0, 0
}

func (i Discount) ToString() string {
	myString, _ := json.Marshal(i)
	return string(myString)
}
