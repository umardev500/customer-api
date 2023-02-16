package usecase

import "customer-api/domain"

type orderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		repository: repository,
	}
}
