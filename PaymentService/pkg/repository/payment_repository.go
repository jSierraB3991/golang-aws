package repository

import (
	"errors"
	"fmt"

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
	ErrorCouldNotDelteuitem    = "could not delete item"
)

type PaymentRepository struct {
	dynaClient dynamodbiface.DynamoDBAPI
	tableName  string
}

func New(session *session.Session) *PaymentRepository {
	return &PaymentRepository{
		tableName:  "payment-table",
		dynaClient: getDynaClient(session),
	}
}

func getDynaClient(awsSession *session.Session) dynamodbiface.DynamoDBAPI {
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
		return nil, errors.New(ErrorCouldNotMarshItem)
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
	repository.DeleteOrder(model)
	_, err = repository.dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("model%v", &model))
	}
	return model, nil
}

func (repository *PaymentRepository) DeleteOrder(model *entity.Payment) error {
	order := model.OrderId
	status := StatusIncomplete
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(repository.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order),
			},
			"status": {
				S: aws.String(status),
			},
		},
	}
	_, err := repository.dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDelteuitem)
	}
	return nil
}
