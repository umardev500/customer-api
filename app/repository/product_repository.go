package repository

import (
	"customer-api/domain"
	"customer-api/pb"
)

type productRepository struct {
	product pb.ProductServiceClient
}

func NewProductRepository(product pb.ProductServiceClient) domain.ProductRepository {
	return &productRepository{
		product: product,
	}
}
