package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

type IDiscountCatalogueRepository repository.IRepository[string, *Discount]

func NewDiscountCatalogue() IDiscountCatalogueRepository {
	return repository.NewMemory[string, *Discount]()
}
