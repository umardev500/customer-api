package delivery

import (
	"customer-api/domain"
	"customer-api/helper"

	"github.com/gofiber/fiber/v2"
)

func (a *appDelivery) UpdateCreds(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	var creds domain.CustomerCreds
	if err := ctx.BodyParser(&creds); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	if err := validate.Struct(&creds); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return nil
}
