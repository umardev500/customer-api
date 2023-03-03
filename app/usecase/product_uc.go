package usecase

import "customer-api/domain"

type productUsecase struct {
	repository domain.ProductRepository
}

func NewProductUsecase(repository domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repository: repository,
	}
}
