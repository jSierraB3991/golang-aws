package repository

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type OrderRepositoty struct {
	dynaClient dynamodbiface.DynamoDBAPI
	tableName  string
}

func New(region string) *OrderRepositoty {
	return &OrderRepositoty{
		tableName:  "order-table",
		dynaClient: getDynaClient(region),
	}
}

func getDynaClient(region string) dynamodbiface.DynamoDBAPI {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil
	}
	return dynamodb.New(awsSession)
}

var (
	ErrorToFiledToFetchRecord  = "failed fetch record"
	ErrorFieldToMarShallRecord = "failed unmarshall record"
	ErrorCouldNotMarshItem     = "could not marshall item"
	ErrorCouldNoDynamoPutItem  = "could not dynamo put iem"
	StatusIncomplete           = "INCOMPLETE"
)

func (repository *OrderRepositoty) FindOrderByUserId(userId string) (*entity.Order, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {S: aws.String(userId)},
			"status":  {S: aws.String(StatusIncomplete)},
		},
		TableName: aws.String(repository.tableName),
	}
	result, err := repository.dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorToFiledToFetchRecord)
	}
	item := new(entity.Order)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFieldToMarShallRecord)
	}

	return item, nil
}
func (repository *OrderRepositoty) SaveFinishOrder(model entity.Order) (*entity.Order, error) {
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
	return &model, nil
}
