package usecase

import (
	"context"
	"customer-api/pb"
)

func (p *productUsecase) GetProduct(ctx context.Context, productId string) (res *pb.ProductFindOneResponse, err error) {
	payload := &pb.ProductFindOneRequest{ProductId: productId}
	res, err = p.repository.GetProduct(ctx, payload)

	return
}
