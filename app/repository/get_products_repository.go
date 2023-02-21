package repository

import (
	"context"
	"customer-api/pb"
)

func (p *productRepository) GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error) {
	res, err = p.product.FindAll(ctx, &pb.ProductFindAllRequest{})

	return
}
