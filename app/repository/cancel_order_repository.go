package repository

import (
	"context"
	"customer-api/pb"
)

func (o *orderRepository) Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error) {
	return o.order.Cancel(ctx, req)
}
