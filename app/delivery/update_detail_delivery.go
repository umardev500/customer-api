package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *appDelivery) Update(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	var detailData domain.CustomerDetail = domain.CustomerDetail{}
	if err := ctx.BodyParser(&detailData); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	if helper.IsZero(detailData) {
		return helper.HandleResponse(ctx, nil, http.StatusNotModified, 0, "Not modified", &pb.OperationResponse{IsAffected: false})
	}

	appCtx := ctx.Context()
	userId := claims["user_id"].(string)

	resp, err := a.appUsecase.Update(appCtx, userId, detailData)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Update detail", resp)
}
