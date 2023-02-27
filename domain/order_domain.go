package domain

import (
	"context"
	"customer-api/pb"
)

type OrderProduct struct {
	ProductId   string `json:"product_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Duration    int64  `json:"duration" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type OrderRequest struct {
	Product OrderProduct        `json:"product"`
	Payment BankTransferRequest `json:"payment"`
}

type OrderPayloadRequest struct {
	ProductId string              `json:"product_id" validate:"required"`
	Payment   BankTransferRequest `json:"payment"`
}

type OrderUsecase interface {
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
	CreateOrder(ctx context.Context, payload OrderRequest, buyer *pb.OrderBuyer, chargeResult BankResponse) (err error)
	Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error)
}
type OrderRepository interface {
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
	CreateOrder(ctx context.Context, req *pb.OrderCreateRequest) (err error)
	Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error)
}
