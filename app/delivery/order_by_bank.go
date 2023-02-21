package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Bank(ctx *fiber.Ctx, orderId string, payload domain.OrderRequest) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	result, err := o.pamyment.ChargeBank(ctx.Context(), orderId, payload.Payment)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	user := claims["user"].(string)
	userId := claims["user_id"].(string)
	name := claims["name"].(string)

	buyer := &pb.OrderBuyer{
		CustomerId: userId,
		User:       user,
		Name:       name,
	}

	// create order
	err = o.usecase.CreateOrder(ctx.Context(), payload, buyer, result)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 201, 0, "Bank payment", result)
}
