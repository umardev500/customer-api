package usecase

import (
	"context"
	"customer-api/domain"
	"encoding/json"
)

func (p *paymentUsecase) ChargeBank(ctx context.Context, orderId string, payment domain.BankTransferRequest) (res domain.BankResponse, err error) {
	payment.TransactionDetails.OrderId = orderId
	payload, _ := json.Marshal(payment)
	res, err = p.repository.ChargeBank(ctx, payload)
	return
}
