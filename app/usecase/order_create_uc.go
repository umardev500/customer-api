package usecase

import (
	"context"
	"customer-api/domain"
	"customer-api/helper"
	"customer-api/pb"
	"customer-api/variable"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (o *orderUsecase) CreateOrder(ctx context.Context, payload domain.OrderRequest, buyer *pb.OrderBuyer, chargeResult domain.BankResponse) (err error) {
	product := &pb.OrderProduct{
		ProductId:   payload.Product.ProductId,
		Name:        payload.Product.Name,
		Price:       payload.Product.Price,
		Duration:    payload.Product.Duration,
		Description: payload.Product.Description,
	}

	va := chargeResult.VaNumbers[0].VaNumber
	bank := chargeResult.VaNumbers[0].Bank
	orderId := chargeResult.OrderID
	amount, _ := strconv.Atoi(strings.Split(chargeResult.GrossAmount, ".")[0])

	payment := &pb.OrderPayment{
		PaymentType: chargeResult.PaymentType,
		OrderId:     orderId,
		Bank:        bank,
		VaNumber:    va,
		GrossAmount: int64(amount),
	}

	loc, err := helper.WIB()
	fmt.Println(err)

	trxTimeStr := chargeResult.TransactionTime
	t, err := time.ParseInLocation(variable.TimeLayout, trxTimeStr, loc)
	if err != nil {
		return
	}
	trxTime := t.UTC().Unix()

	value := &pb.OrderCreateRequest{
		Buyer:   buyer,
		Product: product,
		Payment: payment,
		TrxTime: trxTime,
	}

	err = o.repository.CreateOrder(ctx, value)

	return
}
