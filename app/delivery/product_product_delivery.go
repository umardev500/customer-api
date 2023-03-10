package delivery

import (
	"customer-api/helper"

	"github.com/gofiber/fiber/v2"
)

func (p *productDelivery) GetProduct(ctx *fiber.Ctx) (err error) {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	productId := ctx.Params("id")
	res, err := p.usecase.GetProduct(ctx.Context(), productId)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Find all products", res)
}
