package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

type ICatalogueRepository repository.IRepository[string, IItem]

func NewCatalogue() ICatalogueRepository {
	return repository.NewMemory[string, IItem]()
}
