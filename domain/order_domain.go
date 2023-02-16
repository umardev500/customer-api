package domain

import (
	"context"
	"customer-api/pb"
)

type OrderUsecase interface {
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
}
type OrderRepository interface {
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
}
