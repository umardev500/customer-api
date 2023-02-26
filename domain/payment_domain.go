package domain

import "context"

type BankPaymentResponse struct {
	PaymentType string
	OrderId     string
	Bank        string
	VaNumber    string
	GrossAmount int64
}

type BankTransfer struct {
	Bank     string `json:"bank,omitempty" validate:"required"`
	VaNumber string `json:"va_number,omitempty"`
	Permata  *struct {
		RecipientName string `json:"recipient_name,omitempty"`
	} `json:"permata,omitempty"`
	FreeText *struct {
		Inquiry []struct {
			ID string `json:"id,omitempty"`
			EN string `json:"en,omitempty"`
		} `json:"inquiry,omitempty"`
		Payment []struct {
			ID string `json:"id,omitempty"`
			EN string `json:"en,omitempty"`
		} `json:"payment,omitempty"`
	} `json:"free_text,omitempty"`
	Bca *struct {
		SubCompanyCode string `json:"sub_company_code,omitempty"`
	} `json:"bca,omitempty"`
}

type BankTransferRequest struct {
	PaymentType        string       `json:"payment_type,omitempty" validate:"required"`
	BankTransfer       BankTransfer `json:"bank_transfer,omitempty"`
	TransactionDetails *struct {
		OrderId     string `json:"order_id,omitempty"`
		GrossAmount int64  `json:"gross_amount,omitempty"`
	} `json:"transaction_details,omitempty"`
	CustomerDetails *struct {
		Email     string `json:"email,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Phone     string `json:"phone,omitempty"`
	} `json:"customer_details,omitempty"`
	ItemDetails []struct {
		ID       string `json:"id,omitempty"`
		Price    int    `json:"price,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"item_details,omitempty"`
}

type BankResponse struct {
	StatusCode        string `json:"status_code,omitempty"`
	StatusMessage     string `json:"status_message,omitempty"`
	TransactionID     string `json:"transaction_id,omitempty"`
	OrderID           string `json:"order_id,omitempty"`
	GrossAmount       string `json:"gross_amount,omitempty"`
	PaymentType       string `json:"payment_type,omitempty"`
	TransactionTime   string `json:"transaction_time,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	FraudStatus       string `json:"fraud_status,omitempty"`
	Currency          string `json:"currency,omitempty"`
	VaNumbers         []struct {
		Bank     string `json:"bank,omitempty"`
		VaNumber string `json:"va_number,omitempty"`
	} `json:"va_numbers,omitempty"`
	PermataVaNumber string `json:"permata_va_number,omitempty"`
	MerchantID      string `json:"merchant_id,omitempty"`
}

type PaymentUsecase interface {
	ChargeBank(ctx context.Context, orderId string, payment BankTransferRequest) (res BankResponse, err error)
}

type PaymentRepository interface {
	ChargeBank(ctx context.Context, payload []byte) (res BankResponse, err error)
}
