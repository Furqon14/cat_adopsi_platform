package service

import (
	"cat_adoption_platform/repository"
)

type MidtransService interface {
	Payment() error
}

type midtransService struct {
	// Your implementation here
	repo repository.MidtransRepository
}

func (ms *midtransService) Payment() error {
	// Your implementation here
	return ms.repo.Payment()
}

func NewMidtransService(repo repository.MidtransRepository) MidtransService {
	return &midtransService{repo: repo}
}
