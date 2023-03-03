package delivery

import (
	"customer-api/domain"

	"github.com/gofiber/fiber/v2"
)

type orderDelivery struct {
	usecase domain.OrderUsecase
	product domain.ProductUsecase
	user    domain.AppUsecase
	payment domain.PaymentUsecase
}

func NewOrderDelivery(router fiber.Router, usecase domain.OrderUsecase, payment domain.PaymentUsecase, product domain.ProductUsecase, user domain.AppUsecase) {
	handler := &orderDelivery{
		usecase: usecase,
		payment: payment,
		product: product,
		user:    user,
	}

	router.Get("/orders", handler.Orders)
	router.Get("/orders/:id", handler.GetOrder)
	router.Post("/orders", handler.Create)
	router.Put("/orders/:id/cancel", handler.Cancel)
}
