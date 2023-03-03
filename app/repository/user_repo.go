package repository

import (
	"customer-api/domain"
	"customer-api/pb"
)

type appRepository struct {
	customer pb.CustomerServiceClient
}

func NewAppRepository(customer pb.CustomerServiceClient) domain.AppRepository {
	return &appRepository{
		customer: customer,
	}
}
