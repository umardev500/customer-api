package repository

import (
	"context"
	"customer-api/pb"
)

func (o *orderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = o.order.FindAll(ctx, req)
	return
}
