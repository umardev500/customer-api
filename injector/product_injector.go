package injector

import (
	"customer-api/app/delivery"
	"customer-api/app/repository"
	"customer-api/app/usecase"
	"customer-api/pb"

	"github.com/gofiber/fiber/v2"
)

func NewProductInjector(router fiber.Router, product pb.ProductServiceClient) {
	pr := repository.NewProductRepository(product)
	pu := usecase.NewProductUsecase(pr)
	delivery.NewProductDelivery(router, pu)
}
