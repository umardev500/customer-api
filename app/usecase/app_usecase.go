package usecase

import (
	"context"
	"customer-api/domain"
	"customer-api/pb"
)

type appUsecase struct {
	repository domain.AppRepository
}

func NewAppUsecase(repository domain.AppRepository) domain.AppUsecase {
	return &appUsecase{
		repository: repository,
	}
}

func (a *appUsecase) Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error) {
	res, err = a.repository.Find(ctx, userId)
	return
}
