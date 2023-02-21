package domain

import (
	"context"
	"customer-api/pb"
)

type ProductUsecase interface {
	GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error)
}

type ProductRepository interface {
	GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error)
}
