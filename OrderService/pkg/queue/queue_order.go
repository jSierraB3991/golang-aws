package queue

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jsierrab3991/order-service/pkg/entity"
)

type QueuePayment struct {
	queueName string
	sess      *session.Session
}

func (q *QueuePayment) GetQueueURL() (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(q.sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(q.queueName),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func New(sess *session.Session) *QueuePayment {
	return &QueuePayment{
		queueName: "queue-order-payment",
		sess:      sess,
	}
}

func (q *QueuePayment) SendMsg(queueURL *string, body []byte) error {
	svc := sqs.New(q.sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(body)),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Create Order"),
			},
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Order Service"),
			},
		},
		QueueUrl: queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func (q *QueuePayment) SendMessageQueue(order *entity.Order) error {
	result, err := q.GetQueueURL()
	if err != nil {
		return err
	}

	queueURL := result.QueueUrl
	body, _ := json.Marshal(order)

	err = q.SendMsg(queueURL, body)
	if err != nil {
		return err
	}
	return nil
}
