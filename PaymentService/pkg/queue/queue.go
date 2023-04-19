package queue

import "github.com/jsierrab3991/payment-service/pkg/entity"

type Queue interface {
	SendMessageQueue(*entity.Payment) error
}
