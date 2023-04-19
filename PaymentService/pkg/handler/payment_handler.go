package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jsierrab3991/payment-service/pkg/dto"
	"github.com/jsierrab3991/payment-service/pkg/queue"
	"github.com/jsierrab3991/payment-service/pkg/repository"
	"github.com/jsierrab3991/payment-service/pkg/service"
)

type PaymentHandler struct {
	impl *service.PaymentService
}

func New(region string) *PaymentHandler {
	session := getSession(region)
	return &PaymentHandler{
		impl: service.New(repository.New(session), queue.New(session)),
	}
}

func getSession(region string) *session.Session {
	awsSession, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	return awsSession
}

var (
	ErrorMethodNotAllowed = "method not allowed"
	ErrorInvalidUserData  = "invalid user data"
)

func (handler *PaymentHandler) Payment(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var paymentRequest dto.ProcessPaymentRequest

	if err := json.Unmarshal([]byte(req.Body), &paymentRequest); err != nil {
		return apiResponse(http.StatusBadRequest, ErrorInvalidUserData)
	}
	result, err := handler.impl.PayOrder(paymentRequest)

	if err != nil {
		return apiResponse(http.StatusBadRequest, err.Error())
	}

	return apiResponse(http.StatusCreated, result)
}

func (PaymentHandler) UnHandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
