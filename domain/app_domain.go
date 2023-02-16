package domain

import (
	"context"
	"customer-api/pb"
)

type AppUsecase interface {
	Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error)
}

type AppRepository interface {
	Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error)
}
