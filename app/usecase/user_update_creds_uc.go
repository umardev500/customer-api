package usecase

import (
	"context"
	"customer-api/domain"
	"customer-api/pb"
	"fmt"
)

func (a *appUsecase) UpdateCreds(ctx context.Context, creds domain.CustomerUpdateCredsRequest) (res *pb.OperationResponse, err error) {
	payload := &pb.CustomerUpdateCredsRequest{
		User:    creds.User,
		Pass:    creds.Pass,
		NewPass: creds.NewPass,
	}

	fmt.Println(payload)

	res, err = a.repository.UpdateCreds(ctx, payload)
	return
}
