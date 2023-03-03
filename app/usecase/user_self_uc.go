package usecase

import (
	"context"
	"customer-api/pb"
)

func (a *appUsecase) Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error) {
	res, err = a.repository.Find(ctx, userId)
	return
}
