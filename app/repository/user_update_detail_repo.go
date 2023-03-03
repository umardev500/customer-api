package repository

import (
	"context"
	"customer-api/pb"
)

func (a *appRepository) Update(ctx context.Context, req *pb.CustomerUpdateDetailRequest) (res *pb.OperationResponse, err error) {
	res, err = a.customer.UpdateDetail(ctx, req)

	return
}
