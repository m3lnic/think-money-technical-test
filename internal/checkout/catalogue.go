package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

type ICatalogueRepository interface {
	repository.IRepository[string, *Item]
	GetByItemName(string) (string, *Item, error)
}

func NewCatalogue() ICatalogueRepository {
	return &catalogueRepository{
		IRepository: repository.NewMemory[string, *Item](),
	}
}

type catalogueRepository struct {
	repository.IRepository[string, *Item]
}

func (cr *catalogueRepository) GetByItemName(name string) (string, *Item, error) {
	for key, current := range cr.All() {
		if current.GetName() == name {
			return key, current, nil
		}
	}

	return "", nil, repository.ErrKeyNotFound
}
