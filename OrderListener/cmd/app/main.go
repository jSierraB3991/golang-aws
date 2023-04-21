package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jsierrab3991/order-listener/pkg/repository"
	"github.com/jsierrab3991/order-listener/pkg/service"
)

var region string

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	session := getSession(region)
	os := service.OrderService{
		Repository: repository.New(session),
	}
	for _, message := range sqsEvent.Records {
		log.Printf("id: %v, source :%v = %v", message.MessageId, message.EventSource, message.Body)
		os.Save(message.Body)
	}
	return nil
}

func getSession(region string) *session.Session {
	awsSession, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	return awsSession
}

func main() {
	region = os.Getenv("AWS_REGION")
	lambda.Start(handler)
}
