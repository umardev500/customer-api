package repository

import (
	"context"
	"customer-api/pb"
)

func (o *orderRepository) Find(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.order.FindOne(ctx, req)
	return
}
