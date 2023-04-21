package repository

import "github.com/jsierrab3991/order-listener/pkg/entity"

type Repository interface {
	SaveUpdate(entity *entity.Payment) (*entity.Payment, error)
}
