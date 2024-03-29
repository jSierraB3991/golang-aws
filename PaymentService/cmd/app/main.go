package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	handlers "github.com/jsierrab3991/payment-service/pkg/handler"
)

var region string

func handler(req events.APIGatewayProxyRequest) (response *events.APIGatewayProxyResponse, err error) {
	handlerImpl := handlers.New(region)
	switch req.HTTPMethod {
	case "POST":
		return handlerImpl.Payment(req)
	default:
		return handlerImpl.UnHandledMethod()
	}
}

func main() {
	region = os.Getenv("AWS_REGION")
	_, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return
	}
	lambda.Start(handler)
}
