package repository

import (
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
