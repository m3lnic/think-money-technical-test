package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

func NewCatalogue() repository.IRepository[string, IItem] {
	return repository.NewMemory[string, IItem]()
}
