package usecase

import "customer-api/domain"

type paymentUsecase struct {
	repository domain.PaymentRepository
}

func NewPaymentUsecase(repository domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{
		repository: repository,
	}
}
