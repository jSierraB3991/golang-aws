package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/queue"
	"github.com/jsierrab3991/order-service/pkg/repository"
	"github.com/jsierrab3991/order-service/pkg/service"
)

type OrderHandler struct {
	impl *service.OrderService
}

func New(region string) *OrderHandler {
	session := getSession(region)
	return &OrderHandler{
		impl: service.New(repository.New(session), queue.New(session)),
	}
}

func getSession(region string) *session.Session {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil
	}
	return awsSession
}

var (
	ErrorMethodNotAllowed = "method not allowed"
	ErrorInvalidUserData  = "invalid user data"
)

func (handler *OrderHandler) Order(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var orderRequest dto.CreateOrderRequest

	if err := json.Unmarshal([]byte(req.Body), &orderRequest); err != nil {
		return apiResponse(http.StatusBadRequest, err.Error())
	}

	result, err := handler.impl.CreateOrUpdateOrder(orderRequest)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err.Error())
	}
	return apiResponse(http.StatusCreated, result)
}

func (OrderHandler) UnHandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
