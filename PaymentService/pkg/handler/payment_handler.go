package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jsierrab3991/payment-service/pkg/dto"
	"github.com/jsierrab3991/payment-service/pkg/repository"
	"github.com/jsierrab3991/payment-service/pkg/service"
)

type PaymentHandler struct {
	impl *service.PaymentService
}

func New(region string) *PaymentHandler {
	return &PaymentHandler{
		impl: service.New(repository.New(region)),
	}
}

var (
	ErrorMethodNotAllowed = "method not allowed"
	ErrorInvalidUserData  = "invalid user data"
)

func (handler *PaymentHandler) Order(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var orderRequest dto.ProcessPaymentRequest

	if err := json.Unmarshal([]byte(req.Body), &orderRequest); err != nil {
		return apiResponse(http.StatusBadRequest, ErrorInvalidUserData)
	}

	result, err := handler.impl.PayOrder(orderRequest)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err)
	}
	return apiResponse(http.StatusCreated, result)
}

func (PaymentHandler) UnHandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
