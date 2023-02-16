package injector

import (
	"customer-api/app/delivery"
	"customer-api/app/repository"
	"customer-api/app/usecase"
	"customer-api/pb"

	"github.com/gofiber/fiber/v2"
)

func NewOrderInjector(router fiber.Router, order pb.OrderServiceClient) {
	repo := repository.NewOrderRepository(order)
	uc := usecase.NewOrderUsecase(repo)
	delivery.NewOrderDelivery(router, uc)
}
