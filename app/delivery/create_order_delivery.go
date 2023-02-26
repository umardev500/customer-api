package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Create(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	userId := claims["user_id"].(string)
	res, err := o.user.Find(ctx.Context(), userId)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusInternalServerError, err.Error(), nil)

	}

	if res.IsEmpty {
		err = errors.New("not found")
		return helper.HandleResponse(ctx, err, 0, http.StatusNotFound, http.StatusText(404), res)
	}

	now := time.Now().UTC().Unix()
	expTime := res.Payload.ExpUntil
	isActive := now < expTime
	if isActive {
		err = errors.New("Bad request")
		return helper.HandleResponse(ctx, err, 0, 400, err.Error(), nil)
	}

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
		return o.Bank(ctx, orderId, *payload, claims)
	}

	return ctx.JSON(payload)
}
