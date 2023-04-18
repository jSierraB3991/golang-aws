package repository

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jsierrab3991/payment-service/pkg/entity"
)

var (
	ErrorToFiledToFetchRecord  = "failed fetch record"
	ErrorFieldToMarShallRecord = "failed unmarshall record"
	ErrorCouldNotMarshItem     = "could not marshall item"
	ErrorCouldNoDynamoPutItem  = "could not dynamo put iem"
	StatusIncomplete           = "INCOMPLETE"
	StatusComplete             = "COMPLETE"
)

type PaymentRepository struct {
	dynaClient dynamodbiface.DynamoDBAPI
	tableName  string
}

func New(region string) *PaymentRepository {
	return &PaymentRepository{
		tableName:  "payment-table",
		dynaClient: getDynaClient(region),
	}
}

func getDynaClient(region string) dynamodbiface.DynamoDBAPI {
	awsSession, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	return dynamodb.New(awsSession)
}

func (repository *PaymentRepository) FindOrderByOrderId(orderId string) (*entity.Payment, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {S: aws.String(orderId)},
			"status":   {S: aws.String(StatusIncomplete)},
		},
		TableName: aws.String(repository.tableName),
	}
	result, err := repository.dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorToFiledToFetchRecord)
	}
	item := new(entity.Payment)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFieldToMarShallRecord)
	}

	return item, nil
}
func (repository *PaymentRepository) PaymentOrder(model *entity.Payment) (*entity.Payment, error) {
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshItem)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(repository.tableName),
		Item:      av,
	}
	_, err = repository.dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNoDynamoPutItem)
	}
	return model, nil
}
