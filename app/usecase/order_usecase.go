package usecase

import (
	"context"
	"customer-api/domain"
	"customer-api/pb"
)

type orderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		repository: repository,
	}
}

func (o *orderUsecase) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	return o.repository.FindAll(ctx, req)
}
