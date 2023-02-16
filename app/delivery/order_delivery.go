package delivery

import (
	"customer-api/domain"
	"customer-api/helper"

	"github.com/gofiber/fiber/v2"
)

type orderDelivery struct {
	usecase domain.OrderUsecase
}

func NewOrderDelivery(router fiber.Router, usecase domain.OrderUsecase) {
	handler := &orderDelivery{
		usecase: usecase,
	}

	router.Get("/order-list", handler.Orders)
}

func (o *orderDelivery) Orders(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	return ctx.JSON(claims)
}
