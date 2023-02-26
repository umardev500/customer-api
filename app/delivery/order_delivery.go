package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderDelivery struct {
	usecase  domain.OrderUsecase
	product  domain.ProductUsecase
	pamyment domain.PaymentUsecase
}

func NewOrderDelivery(router fiber.Router, usecase domain.OrderUsecase, payment domain.PaymentUsecase, product domain.ProductUsecase) {
	handler := &orderDelivery{
		usecase:  usecase,
		pamyment: payment,
		product:  product,
	}

	router.Get("/orders", handler.Orders)
	router.Post("/orders", handler.Create)
}

func (o *orderDelivery) Orders(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	appCtx := ctx.Context()
	userId := claims["user_id"].(string)

	sort := ctx.Query("sort", "asc")
	page := helper.ToInt(ctx.Query("page", "0"))
	perPage := helper.ToInt(ctx.Query("per_page"))
	search := ctx.Query("search")
	status := ctx.Query("status")
	countOnly, _ := strconv.ParseBool(ctx.Query("count_only"))
	filter := &pb.OrderFindAllRequest{
		Sort:      sort,
		Page:      page,
		PerPage:   perPage,
		Search:    search,
		Status:    status,
		CountOnly: countOnly,
		UserId:    userId,
	}
	res, err := o.usecase.FindAll(appCtx, filter)

	return helper.HandleResponse(ctx, err, 200, 0, "Find orders", res)
}
