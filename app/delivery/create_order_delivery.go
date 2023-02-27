package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
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

	isDeleted := res.Payload.DeletedAt != 0
	if isDeleted {
		err = errors.New("bad request")
		return helper.HandleResponse(ctx, err, 0, 400, err.Error(), nil)
	}

	now := time.Now().UTC().Unix()
	expTime := res.Payload.ExpUntil
	isActive := now < expTime
	if isActive {
		err = errors.New("bad request")
		return helper.HandleResponse(ctx, err, 0, 400, err.Error(), nil)
	}

	buyer := &pb.OrderBuyer{
		CustomerId: userId,
		User:       res.Payload.User,
		Name:       res.Payload.Detail.Name,
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
		return o.Bank(ctx, orderId, *payload, buyer)
	}

	return ctx.JSON(payload)
}
