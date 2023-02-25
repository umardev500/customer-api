package usecase

import (
	"context"
	"customer-api/domain"
	"customer-api/pb"
)

func (a *appUsecase) Update(ctx context.Context, userId string, detail domain.CustomerDetail) (res *pb.OperationResponse, err error) {
	var locationPayload *pb.CustomerLocation

	if detail.Location != nil {
		locationPayload = &pb.CustomerLocation{
			Address:    detail.Location.Address,
			Village:    detail.Location.Village,
			District:   detail.Location.District,
			City:       detail.Location.City,
			Province:   detail.Location.Province,
			PostalCode: detail.Location.PostalCode,
		}
	}

	detailPayload := &pb.CustomerDetail{
		Npsn:     detail.Npsn,
		Name:     detail.Name,
		Email:    detail.Email,
		Wa:       detail.Wa,
		Type:     detail.Type,
		Level:    detail.Level,
		About:    detail.About,
		Logo:     detail.Logo,
		Location: locationPayload,
	}

	payload := &pb.CustomerUpdateDetailRequest{
		CustomerId: userId,
		Detail:     detailPayload,
	}

	res, err = a.repository.Update(ctx, payload)

	return
}
