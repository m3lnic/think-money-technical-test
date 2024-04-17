package checkout

import "github.com/m3lnic/think-money-technical-test/internal/repository"

type IDiscountCatalogueRepository interface {
	repository.IRepository[string, *Discount]
}

func NewDiscountCatalogue(catalogue ICatalogueRepository) IDiscountCatalogueRepository {
	return &discountCatalogueRepository{
		IRepository: repository.NewMemory[string, *Discount](),
		catalogue:   catalogue,
	}
}

type discountCatalogueRepository struct {
	repository.IRepository[string, *Discount]
	catalogue ICatalogueRepository
}

func (dcr *discountCatalogueRepository) Create(key string, data *Discount) (*Discount, error) {
	if _, err := dcr.catalogue.Read(key); err != nil {
		return nil, err
	}

	return dcr.IRepository.Create(key, data)
}
