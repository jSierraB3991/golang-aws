package service

import (
	"github.com/google/uuid"
	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/entity"
	"github.com/jsierrab3991/order-service/pkg/repository"
)

type OrderService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

var (
	StatusIncomplete = "INCOMPLETE"
)

func (impl *OrderService) CreateOrUpdateOrder(orderRequest dto.CreateOrderRequest) (*entity.Order, error) {
	item, err := impl.repo.FindOrderByUserId(orderRequest.UserID)
	if err != nil {
		return nil, err
	}
	if item.UserID == "" || item.OrderID == "" {
		return impl.createNewOrder(orderRequest)
	}
	return impl.updateOrder(item, orderRequest)
}

func (impl *OrderService) createNewOrder(request dto.CreateOrderRequest) (*entity.Order, error) {

	model := requestToOrder(request)
	return impl.repo.SaveFinishOrder(model)
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

func requestToOrder(request dto.CreateOrderRequest) entity.Order {
	return entity.Order{
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
}
