package service

import (
	"github.com/google/uuid"
	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/entity"
	"github.com/jsierrab3991/order-service/pkg/queue"
	"github.com/jsierrab3991/order-service/pkg/repository"
)

type OrderService struct {
	repo  repository.Repository
	queue queue.Queue
}

func New(repo repository.Repository, queue queue.Queue) *OrderService {
	return &OrderService{
		repo:  repo,
		queue: queue,
	}
}

var (
	StatusIncomplete = "INCOMPLETE"
)

func (impl *OrderService) CreateOrUpdateOrder(orderRequest dto.CreateOrderRequest) (*entity.Order, error) {
	item, err := impl.repo.FindOrderByUserId(orderRequest.UserID)
	var response *entity.Order
	if err != nil || item.OrderID == "" {
		response, err = impl.createNewOrder(orderRequest)
	} else {
		response, err = impl.updateOrder(item, orderRequest)
	}

	if err != nil {
		return nil, err
	}
	if err = impl.queue.SendMessageQueue(response); err != nil {
		return nil, err
	}
	return response, nil
}

func (impl *OrderService) createNewOrder(request dto.CreateOrderRequest) (*entity.Order, error) {
	model := &entity.Order{
		OrderID:    uuid.NewString(),
		UserID:     request.UserID,
		TotalPrice: request.TotalPrice,
		Status:     StatusIncomplete,
		List: []entity.OrderDetail{
			{
				Item:       request.Item,
				Quantity:   request.Quantity,
				TotalPrice: request.TotalPrice,
			},
		},
	}
	return impl.repo.SaveFinishOrder(*model)
}

func (impl *OrderService) updateOrder(item *entity.Order, request dto.CreateOrderRequest) (*entity.Order, error) {
	model := requestToDetail(request)
	item.List = append(item.List, model)
	item.TotalPrice += model.TotalPrice
	return impl.repo.SaveFinishOrder(*item)
}

func requestToDetail(request dto.CreateOrderRequest) entity.OrderDetail {
	return entity.OrderDetail{
		Item:       request.Item,
		Quantity:   request.Quantity,
		TotalPrice: request.TotalPrice,
	}
}
