package checkout

import (
	"github.com/m3lnic/think-money-technical-test/internal/repository"
)

type ICheckout interface {
	Scan(string) error
	GetTotal() (int, error)
}

func New(catalogue ICatalogueRepository, discountCatalogue IDiscountCatalogueRepository) ICheckout {
	return &checkout{
		catalogue:         catalogue,
		discountCatalogue: discountCatalogue,
		scannedItems:      repository.NewMemory[string, int](),
	}
}

type checkout struct {
	catalogue         ICatalogueRepository
	discountCatalogue IDiscountCatalogueRepository
	scannedItems      repository.IRepository[string, int] // > SKU -> Quantity
}

// GetTotal implements ICheckout.
func (c *checkout) GetTotal() (int, error) {
	curTotal := 0

	for sku, currentItemCount := range c.scannedItems.All() {
		catalogueItem, err := c.catalogue.Read(sku)
		if err != nil {
			return 0, err
		}

		totalItems := currentItemCount
		if itemDiscount, err := c.discountCatalogue.Read(sku); err == nil {
			discountTotal, remainingItems := itemDiscount.QualifiesFor(currentItemCount)

			curTotal += discountTotal
			totalItems = remainingItems
		}

		curTotal += totalItems * catalogueItem.GetUnitPrice()
	}

	return curTotal, nil
}

// Scan implements ICheckout.
func (c *checkout) Scan(sku string) error {
	// > If we already have the item, skip the other steps and just add 1
	if myItemCount, err := c.scannedItems.Read(sku); err == nil {
		c.scannedItems.Update(sku, myItemCount+1)
		return nil
	}

	// > Check SKU is valid
	_, err := c.catalogue.Read(sku)
	if err != nil {
		return err
	}

	// > Add to cart
	c.scannedItems.Create(sku, 1)

	return nil
}
