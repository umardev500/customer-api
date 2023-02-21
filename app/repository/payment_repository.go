package repository

import "customer-api/domain"

type paymentRepository struct{}

func NewPaymentRepository() domain.PaymentRepository {
	return &paymentRepository{}
}
