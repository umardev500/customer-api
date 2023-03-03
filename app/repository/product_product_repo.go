package repository

import (
	"context"
	"customer-api/pb"
)

func (p *productRepository) GetProduct(ctx context.Context, req *pb.ProductFindOneRequest) (res *pb.ProductFindOneResponse, err error) {
	return p.product.FindOne(ctx, req)
}
