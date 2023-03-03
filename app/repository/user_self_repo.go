package repository

import (
	"context"
	"customer-api/pb"
)

func (a *appRepository) Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error) {
	res, err = a.customer.Find(ctx, &pb.CustomerFindRequest{CustomerId: userId})

	return
}
