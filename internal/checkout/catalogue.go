package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

type ICatalogueRepository repository.IRepository[string, *Item]

func NewCatalogue() ICatalogueRepository {
	return repository.NewMemory[string, *Item]()
}
