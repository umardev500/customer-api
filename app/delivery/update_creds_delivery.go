package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *appDelivery) UpdateCreds(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	var creds domain.CustomerUpdateCredsRequest
	if err := ctx.BodyParser(&creds); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	user := claims["user"].(string)
	creds.User = user

	if err := validate.Struct(&creds); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	appCtx := ctx.Context()
	resp, err := a.appUsecase.UpdateCreds(appCtx, creds)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	if !resp.IsAffected {
		return helper.HandleResponse(ctx, nil, http.StatusNotModified, 0, "Not modified", resp)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Update detail", resp)
}
