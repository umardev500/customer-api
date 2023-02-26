package injector

import (
	"customer-api/app/delivery"
	"customer-api/app/repository"
	"customer-api/app/usecase"
	"customer-api/domain"
	"customer-api/pb"

	"github.com/gofiber/fiber/v2"
)

func NewOrderInjector(router fiber.Router, order pb.OrderServiceClient, productUsecase domain.ProductUsecase, user domain.AppUsecase) {
	repo := repository.NewOrderRepository(order)
	uc := usecase.NewOrderUsecase(repo)
	paymentRepo := repository.NewPaymentRepository()
	paymentUc := usecase.NewPaymentUsecase(paymentRepo)
	delivery.NewOrderDelivery(router, uc, paymentUc, productUsecase, user)
}
