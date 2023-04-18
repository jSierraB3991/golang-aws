package repository

import "github.com/jsierrab3991/payment-service/pkg/entity"

type Repository interface {
	FindOrderByOrderId(orderId string) (*entity.Payment, error)
	PaymentOrder(model *entity.Payment) (*entity.Payment, error)
}
