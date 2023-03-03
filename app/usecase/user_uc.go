package usecase

import (
	"customer-api/domain"
)

type appUsecase struct {
	repository domain.AppRepository
}

func NewAppUsecase(repository domain.AppRepository) domain.AppUsecase {
	return &appUsecase{
		repository: repository,
	}
}
