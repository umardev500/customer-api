package domain

import (
	"context"
	"customer-api/pb"
)

type ProductUsecase interface {
	GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error)
	GetProduct(ctx context.Context, productId string) (res *pb.ProductFindOneResponse, err error)
}

type ProductRepository interface {
	GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error)
	GetProduct(ctx context.Context, req *pb.ProductFindOneRequest) (res *pb.ProductFindOneResponse, err error)
}
