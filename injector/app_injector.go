package injector

import (
	"customer-api/app/delivery"
	"customer-api/app/repository"
	"customer-api/app/usecase"
	"customer-api/pb"

	"github.com/gofiber/fiber/v2"
)

func NewAppInjector(router fiber.Router, customer pb.CustomerServiceClient) {
	appRepo := repository.NewAppRepository(customer)
	appUsecase := usecase.NewAppUsecase(appRepo)
	delivery.NewAppDelivery(router, appUsecase)
}
