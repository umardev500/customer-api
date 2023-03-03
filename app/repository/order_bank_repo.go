package repository

import (
	"bytes"
	"context"
	"customer-api/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (p *paymentRepository) ChargeBank(ctx context.Context, payload []byte) (res domain.BankResponse, err error) {
	data := bytes.NewBuffer(payload)

	client := http.Client{}
	target := fmt.Sprintf("%s/%s", os.Getenv("PAYMENT_URL"), "charge")
	req, err := http.NewRequest("POST", target, data)
	if err != nil {
		return
	}

	key := os.Getenv("MIDTRANS_KEY")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}
