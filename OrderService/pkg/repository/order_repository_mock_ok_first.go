package repository

import (
	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type OrderRepositoryMockOkFirst struct {
}

func (OrderRepositoryMockOkFirst) FindOrderByUserId(userId string) (*entity.Order, error) {
	return &entity.Order{}, nil
}
func (OrderRepositoryMockOkFirst) SaveFinishOrder(model entity.Order) (*dto.CreateOderEvent, error) {
	return &dto.CreateOderEvent{
		OrderID:    model.OrderID,
		TotalPrice: model.TotalPrice,
	}, nil
}
