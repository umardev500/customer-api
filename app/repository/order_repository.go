package repository

import (
	"context"
	"customer-api/domain"
	"customer-api/pb"
)

type orderRepository struct {
	order pb.OrderServiceClient
}

func NewOrderRepository(order pb.OrderServiceClient) domain.OrderRepository {
	return &orderRepository{
		order: order,
	}
}

func (o *orderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	return o.order.FindAll(ctx, req)
}
