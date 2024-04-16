package checkout

import "encoding/json"

type IItem interface {
	GetName() string
	GetUnitPrice() int
	ToString() string
}

// > Function means that we must always pass these values when creating an item
func NewItem(name string, unitPrice int) IItem {
	return &item{
		Name:      name,
		UnitPrice: unitPrice,
	}
}

type item struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unitPrice"`
}

// GetName implements IItem.
func (i *item) GetName() string {
	return i.Name
}

// GetUnitPrice implements IItem.
func (i *item) GetUnitPrice() int {
	return i.UnitPrice
}

func (i item) ToString() string {
	myString, _ := json.Marshal(i)
	return string(myString)
}
