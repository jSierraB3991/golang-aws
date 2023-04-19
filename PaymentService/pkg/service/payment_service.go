package service

import (
	"github.com/jsierrab3991/payment-service/pkg/dto"
	"github.com/jsierrab3991/payment-service/pkg/entity"
	"github.com/jsierrab3991/payment-service/pkg/queue"
	"github.com/jsierrab3991/payment-service/pkg/repository"
)

type PaymentService struct {
	repo  repository.Repository
	queue queue.Queue
}

func New(repo repository.Repository, queue queue.Queue) *PaymentService {
	return &PaymentService{
		repo:  repo,
		queue: queue,
	}
}

func (impl *PaymentService) PayOrder(request dto.ProcessPaymentRequest) (*entity.Payment, error) {
	item, err := impl.repo.FindOrderByOrderId(request.OrderID)
	if err != nil {
		return nil, err
	}
	item.Status = request.Status
	response, err := impl.repo.PaymentOrder(item)

	if err != nil {
		return nil, err
	}
	if err = impl.queue.SendMessageQueue(response); err != nil {
		return nil, err
	}
	return response, nil
}
