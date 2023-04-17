package repository

import (
	"github.com/google/uuid"
	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type OrderRepositoryMockOkOther struct {
}

func (OrderRepositoryMockOkOther) FindOrderByUserId(userId string) (*entity.Order, error) {
	return dataMock(userId), nil
}
func (OrderRepositoryMockOkOther) SaveFinishOrder(model entity.Order) (*dto.CreateOderEvent, error) {
	data := dataMock(model.UserID)
	return &dto.CreateOderEvent{
		OrderID:    data.OrderID,
		TotalPrice: data.TotalPrice + model.TotalPrice,
	}, nil
}

func dataMock(userId string) *entity.Order {
	return &entity.Order{
		OrderID:    uuid.NewString(),
		UserID:     userId,
		TotalPrice: 2000,
		Status:     StatusIncomplete,
		List: []entity.OrderDetail{{
			Item:       "Alfombra",
			Quantity:   2,
			TotalPrice: 10000,
		}},
	}
}
