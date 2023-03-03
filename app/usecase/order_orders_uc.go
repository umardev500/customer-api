package usecase

import (
	"context"
	"customer-api/pb"
)

func (o *orderUsecase) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	return o.repository.FindAll(ctx, req)
}
