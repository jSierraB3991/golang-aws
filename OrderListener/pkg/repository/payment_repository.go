package repository

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jsierrab3991/order-listener/pkg/entity"
)

var (
	ErrorToFiledToFetchRecord = "failed fetch record"
	ErrorCouldNotMarshItem    = "could not marshall item"
	StatusIncomplete          = "INCOMPLETE"
	ErrorCouldNotDelteitem    = "could not delete item"
)

type PaymentRepository struct {
	dynaClient dynamodbiface.DynamoDBAPI
	tableName  string
}

func getDynaClient(awsSession *session.Session) dynamodbiface.DynamoDBAPI {
	return dynamodb.New(awsSession)
}

func New(session *session.Session) *PaymentRepository {
	return &PaymentRepository{
		tableName:  "payment-table",
		dynaClient: getDynaClient(session),
	}
}

func (rp *PaymentRepository) SaveUpdate(entity *entity.Payment) (*entity.Payment, error) {
	response, err := rp.FindOrderByOrderId(entity.OrderId)
	if err != nil {
		if err.Error() == ErrorToFiledToFetchRecord {
			return rp.saveOrder(entity)
		}
	}
	return rp.updateOrder(response, entity)
}

func (rp *PaymentRepository) saveOrder(model *entity.Payment) (*entity.Payment, error) {
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshItem)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(rp.tableName),
		Item:      av,
	}
	_, err = rp.dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("model %v", *model))
	}
	return model, nil
}

func (rp *PaymentRepository) updateOrder(now *entity.Payment, model *entity.Payment) (*entity.Payment, error) {
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshItem)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(rp.tableName),
		Item:      av,
	}
	rp.deleteOrder(now)
	_, err = rp.dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("model%v", &model))
	}
	return model, nil
}

func (rp *PaymentRepository) FindOrderByOrderId(orderId string) (*entity.Payment, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {S: aws.String(orderId)},
			"status":   {S: aws.String(StatusIncomplete)},
		},
		TableName: aws.String(rp.tableName),
	}
	result, err := rp.dynaClient.GetItem(input)
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

func (rp *PaymentRepository) deleteOrder(model *entity.Payment) error {
	order := model.OrderId
	status := StatusIncomplete
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(rp.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order),
			},
			"status": {
				S: aws.String(status),
			},
		},
	}
	_, err := rp.dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDelteitem)
	}
	return nil
}
