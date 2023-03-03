package delivery

import (
	"customer-api/helper"
	"customer-api/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
