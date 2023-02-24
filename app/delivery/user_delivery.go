package delivery

import (
	"customer-api/domain"
	"customer-api/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type appDelivery struct {
	appUsecase domain.AppUsecase
}

var validate = validator.New()

func NewAppDelivery(router fiber.Router, appUsecase domain.AppUsecase) {
	handler := &appDelivery{
		appUsecase: appUsecase,
	}

	router.Get("/me", handler.Find)
	router.Put("/update-detail", handler.Update)
	router.Put("/update-creds", handler.UpdateCreds)
	router.Put("/logo", handler.UploadLogo)
}

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
