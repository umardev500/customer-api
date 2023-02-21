package usecase

import (
	"context"
	"customer-api/pb"
)

func (p *productUsecase) GetProducts(ctx context.Context) (res *pb.ProductFindAllResponse, err error) {
	res, err = p.repository.GetProducts(ctx)
	return
}
