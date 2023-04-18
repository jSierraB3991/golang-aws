package service

import (
	"github.com/jsierrab3991/payment-service/pkg/dto"
	"github.com/jsierrab3991/payment-service/pkg/entity"
	"github.com/jsierrab3991/payment-service/pkg/repository"
)

type PaymentService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (impl *PaymentService) PayOrder(request dto.ProcessPaymentRequest) (*entity.Payment, error) {
	item, err := impl.repo.FindOrderByOrderId(request.OrderID)
	if err != nil {
		return nil, err
	}
	item.Status = request.Status
	return impl.repo.PaymentOrder(item)
}
