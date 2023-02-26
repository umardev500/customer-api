package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Create(ctx *fiber.Ctx) error {
	paymentType := ctx.Query("payment_type")

	var payload = new(domain.OrderPayloadRequest)
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	if err := validate.Struct(payload); err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	orderId := strconv.Itoa(int(time.Now().UTC().UnixNano()))

	if paymentType == "bank" {
		return o.Bank(ctx, orderId, *payload)
	}

	return ctx.JSON(payload)
}
