package queue

import "github.com/jsierrab3991/order-service/pkg/entity"

type Queue interface {
	SendMessageQueue(order *entity.Order) error
}
