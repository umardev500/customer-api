package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Bank(ctx *fiber.Ctx, orderId string, payload domain.OrderPayloadRequest) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	product, err := o.product.GetProduct(ctx.Context(), payload.ProductId)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	if product.IsEmpty {
		return helper.HandleResponse(ctx, err, 404, http.StatusNotFound, "not found", product)
	}

	orderData := domain.OrderRequest{
		Product: domain.OrderProduct{
			ProductId:   payload.ProductId,
			Name:        product.Payload.Name,
			Price:       product.Payload.Price,
			Duration:    product.Payload.Duration,
			Description: product.Payload.Description,
		},
		Payment: payload.Payment,
	}

	orderData.Payment.TransactionDetails = &struct {
		OrderId     string `json:"order_id,omitempty"`
		GrossAmount int64  `json:"gross_amount,omitempty"`
	}{
		OrderId:     orderId,
		GrossAmount: product.Payload.Price,
	}

	result, err := o.pamyment.ChargeBank(ctx.Context(), orderId, orderData.Payment)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
	}

	user := claims["user"].(string)
	userId := claims["user_id"].(string)
	name := claims["name"].(string)

	buyer := &pb.OrderBuyer{
		CustomerId: userId,
		User:       user,
		Name:       name,
	}

	if result.VaNumbers == nil {
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return helper.HandleResponse(ctx, err, 0, 400, http.StatusText(http.StatusBadRequest), nil)
	}

	// create order
	err = o.usecase.CreateOrder(ctx.Context(), orderData, buyer, result)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 201, 0, "Bank payment", result)
}
