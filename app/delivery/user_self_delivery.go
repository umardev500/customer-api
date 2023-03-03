package delivery

import (
	"customer-api/helper"

	"github.com/gofiber/fiber/v2"
)

func (a *appDelivery) Find(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	appCtx := ctx.Context()
	userId := claims["user_id"].(string)
	res, err := a.appUsecase.Find(appCtx, userId)
	return helper.HandleResponse(ctx, err, 200, 0, "Find customer", res)
}
