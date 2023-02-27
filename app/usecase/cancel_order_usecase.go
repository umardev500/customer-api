package usecase

import (
	"context"
	"customer-api/pb"
)

func (o *orderUsecase) Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error) {
	return o.repository.Cancel(ctx, req)
}
