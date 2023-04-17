package repository

import (
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type Repository interface {
	FindOrderByUserId(userId string) (*entity.Order, error)
	SaveFinishOrder(model entity.Order) (*entity.Order, error)
}
