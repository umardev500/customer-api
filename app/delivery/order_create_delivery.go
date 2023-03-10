package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"customer-api/variable"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Create(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	// check for previous order not paid
	orders, err := o.usecase.FindAll(ctx.Context(), &pb.OrderFindAllRequest{Status: "pending"})
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusInternalServerError, err.Error(), nil)
	}

	if !orders.IsEmpty {
		// o01 for prevoius order not ended
		err = errors.New("o01")
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	// find user
	userId := claims["user_id"].(string)
	res, err := o.user.Find(ctx.Context(), userId)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusInternalServerError, err.Error(), nil)
	}

	if res.IsEmpty {
		// u00 for user not found
		err = errors.New("u00")
		return helper.HandleResponse(ctx, err, 0, http.StatusNotFound, http.StatusText(404), res)
	}

	isDeleted := res.Payload.DeletedAt != 0
	if isDeleted {
		// u01 for deleted user
		err = errors.New("u01")
		return helper.HandleResponse(ctx, err, 0, 400, err.Error(), nil)
	}

	now := time.Now().UTC().Unix()
	expTime := res.Payload.ExpUntil
	isActive := now < expTime
	if isActive {
		// r01 for rent still active
		err = errors.New("r01")
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

	// set expiration
	payExp := helper.ToInt(os.Getenv("PAY_EXP"))
	payload.Payment.CustomExpiry = &struct {
		ExpiryDuration int    "json:\"expiry_duration\""
		Unit           string "json:\"unit\""
	}{
		ExpiryDuration: int(payExp),
		Unit:           "minute",
	}

	// validate struct
	if err := validate.Struct(payload); err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	// create order id
	orderId := helper.ToString(int(helper.GetTime(&variable.UnixNano)))

	if paymentType == "bank" {
		return o.Bank(ctx, orderId, *payload, buyer)
	}

	return ctx.JSON(payload)
}
