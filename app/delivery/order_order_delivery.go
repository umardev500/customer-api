package delivery

import (
	"customer-api/helper"
	"customer-api/pb"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) GetOrder(ctx *fiber.Ctx) (err error) {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	orderId := ctx.Params("id")
	res, err := o.usecase.Find(ctx.Context(), &pb.OrderFindOneRequest{OrderId: orderId})

	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Find all products", res)
}
