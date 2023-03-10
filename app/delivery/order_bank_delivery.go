package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (o *orderDelivery) Bank(ctx *fiber.Ctx, orderId string, payload domain.OrderPayloadRequest, buyer *pb.OrderBuyer) error {
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

	result, err := o.payment.ChargeBank(ctx.Context(), orderId, orderData.Payment)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, http.StatusBadRequest, err.Error(), nil)
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
