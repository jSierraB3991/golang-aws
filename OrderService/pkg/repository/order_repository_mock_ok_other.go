package repository

import (
	"github.com/google/uuid"
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type OrderRepositoryMockOkOther struct {
	PreValue int64
}

func (orm OrderRepositoryMockOkOther) FindOrderByUserId(userId string) (*entity.Order, error) {
	return dataMock(userId, orm.PreValue), nil
}
func (orm OrderRepositoryMockOkOther) SaveFinishOrder(model entity.Order) (*entity.Order, error) {
	return &model, nil
}

func dataMock(userId string, prevalue int64) *entity.Order {
	return &entity.Order{
		OrderID:    uuid.NewString(),
		UserID:     userId,
		TotalPrice: prevalue,
		Status:     StatusIncomplete,
		List: []entity.OrderDetail{{
			Item:       "Alfombra",
			Quantity:   2,
			TotalPrice: prevalue,
		}},
	}
}
