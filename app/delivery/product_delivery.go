package delivery

import (
	"customer-api/domain"

	"github.com/gofiber/fiber/v2"
)

type productDelivery struct {
	usecase domain.ProductUsecase
}

func NewProductDelivery(router fiber.Router, usecase domain.ProductUsecase) {
	handler := &productDelivery{
		usecase: usecase,
	}
	router.Get("/products", handler.GetProducts)
	router.Get("/products/:id", handler.GetProduct)
}
