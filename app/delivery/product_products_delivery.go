package delivery

import (
	"customer-api/helper"

	"github.com/gofiber/fiber/v2"
)

func (p productDelivery) GetProducts(ctx *fiber.Ctx) error {
	appCtx := ctx.Context()
	resp, err := p.usecase.GetProducts(appCtx)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), resp)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Find all products", resp)
}
