package domain

import (
	"context"
	"customer-api/pb"
)

type CustomerLocation struct {
	Address    string `json:"address,omitempty"`
	Village    string `json:"village,omitempty"`
	District   string `json:"district,omitempty"`
	City       string `json:"city,omitempty"`
	Province   string `json:"province,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type CustomerDetail struct {
	Npsn     string            `json:"npsn,omitempty"`
	Name     string            `json:"name,omitempty"`
	Email    string            `json:"email,omitempty"`
	Wa       string            `json:"wa,omitempty"`
	Type     string            `json:"type,omitempty"`
	Level    string            `json:"level,omitempty"`
	About    string            `json:"about,omitempty"`
	Logo     string            `json:"logo,omitempty"`
	Location *CustomerLocation `json:"location,omitempty"`
}

type Customer struct {
	CustomerId string          `json:"customer_id,omitempty"`
	User       string          `json:"user,omitempty"`
	Pass       string          `json:"pass,omitempty"`
	Detail     *CustomerDetail `json:"detail,omitempty"`
	Status     string          `json:"status,omitempty"`
	ExpUntil   int64           `json:"exp_until,omitempty"`
	CreatedAt  int64           `json:"created_at,omitempty"`
	UpdatedAt  int64           `json:"updated_at,omitempty"`
	DeletedAt  int64           `json:"deleted_at,omitempty"`
}

type CustomerUpdateCredsRequest struct {
	User    string `validate:"required,min=6"`
	Pass    string `json:"pass" validate:"required,min=6"`
	NewPass string `json:"new_pass" validate:"required,min=6"`
}

type AppUsecase interface {
	Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error)
	Update(ctx context.Context, userId string, detail CustomerDetail) (res *pb.OperationResponse, err error)
	UpdateCreds(ctx context.Context, creds CustomerUpdateCredsRequest) (res *pb.OperationResponse, err error)
}

type AppRepository interface {
	Find(ctx context.Context, userId string) (res *pb.CustomerFindResponse, err error)
	Update(ctx context.Context, req *pb.CustomerUpdateDetailRequest) (res *pb.OperationResponse, err error)
	UpdateCreds(ctx context.Context, req *pb.CustomerUpdateCredsRequest) (res *pb.OperationResponse, err error)
}
