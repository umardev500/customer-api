package delivery

import (
	"customer-api/domain"

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
