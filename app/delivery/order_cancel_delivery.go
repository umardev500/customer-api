package delivery

import (
	"customer-api/helper"
	"customer-api/pb"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Cancel(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	appCtx := ctx.Context()
	orderId := ctx.Params("id")
	userId := claims["user_id"].(string)
	payload := &pb.OrderCancelRequest{
		OrderId: orderId,
		UserId:  userId,
	}

	res, err := o.usecase.Cancel(appCtx, payload)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "cancel order", res)
}
