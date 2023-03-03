package usecase

import (
	"context"
	"customer-api/pb"
)

func (o *orderUsecase) Find(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.repository.Find(ctx, req)
	return
}
