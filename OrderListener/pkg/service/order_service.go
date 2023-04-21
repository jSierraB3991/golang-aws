package service

import (
	"encoding/json"
	"log"

	"github.com/jsierrab3991/order-listener/pkg/dto"
	"github.com/jsierrab3991/order-listener/pkg/entity"
	"github.com/jsierrab3991/order-listener/pkg/repository"
)

type OrderService struct {
	Repository repository.Repository
}

func (os *OrderService) Save(message string) {

	var order dto.Order
	err := json.Unmarshal([]byte(message), &order)
	if err != nil {
		log.Printf("Error un marshall type %v", err.Error())
	}
	response, err := os.Repository.SaveUpdate(os.dtoToModel(order))
	if err != nil {
		log.Printf("fail to save order in payment %v", err.Error())
	}
	log.Printf("sucess save order %v", *response)
}

func (OrderService) dtoToModel(order dto.Order) *entity.Payment {
	return &entity.Payment{
		OrderId:    order.OrderID,
		UserId:     order.UserID,
		Status:     order.Status,
		TotalPrice: order.TotalPrice,
	}
}
