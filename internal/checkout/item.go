package checkout

import "encoding/json"

type IItem interface {
	GetName() string
	GetUnitPrice() int
	ToString() string
}

// > Function means that we must always pass these values when creating an item
func NewItem(name string, unitPrice int) *Item {
	return &Item{
		Name:      name,
		UnitPrice: unitPrice,
	}
}

type Item struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unitPrice"`
}

// GetName implements IItem.
func (i *Item) GetName() string {
	return i.Name
}

// GetUnitPrice implements IItem.
func (i *Item) GetUnitPrice() int {
	return i.UnitPrice
}

func (i Item) ToString() string {
	myString, _ := json.Marshal(i)
	return string(myString)
}
