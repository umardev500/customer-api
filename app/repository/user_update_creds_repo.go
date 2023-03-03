package repository

import (
	"context"
	"customer-api/pb"
)

func (a *appRepository) UpdateCreds(ctx context.Context, req *pb.CustomerUpdateCredsRequest) (res *pb.OperationResponse, err error) {
	return a.customer.UpdateCreds(ctx, req)
}
