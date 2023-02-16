package repository

import (
	"context"
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

func (a *appRepository) Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error) {
	res, err = a.customer.Find(ctx, &pb.CustomerFindRequest{CustomerId: userId})

	return
}
