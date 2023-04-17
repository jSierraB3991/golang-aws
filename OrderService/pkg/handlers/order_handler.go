package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/repository"
	"github.com/jsierrab3991/order-service/pkg/service"
)

type OrderHandler struct {
	impl *service.OrderService
}

func New(region string) *OrderHandler {
	return &OrderHandler{
		impl: service.New(repository.New(region)),
	}
}

var (
	ErrorMethodNotAllowed = "method not allowed"
	ErrorInvalidUserData  = "invalid user data"
)

func (handler *OrderHandler) Order(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var orderRequest dto.CreateOrderRequest

	if err := json.Unmarshal([]byte(req.Body), &orderRequest); err != nil {
		return apiResponse(http.StatusBadRequest, ErrorInvalidUserData)
	}

	result, err := handler.impl.CreateOrUpdateOrder(orderRequest)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err)
	}
	return apiResponse(http.StatusCreated, result)
}

func (OrderHandler) UnHandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
