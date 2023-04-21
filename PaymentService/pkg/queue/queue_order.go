package queue

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jsierrab3991/payment-service/pkg/entity"
)

type QueueOrder struct {
	queueName string
	sess      *session.Session
}

func (q *QueueOrder) GetQueueURL() (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(q.sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(q.queueName),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func New(sess *session.Session) *QueueOrder {
	return &QueueOrder{
		queueName: "queue-payment-complete",
		sess:      sess,
	}
}

func (q *QueueOrder) SendMsg(queueURL *string, body []byte) error {
	svc := sqs.New(q.sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(body)),
		QueueUrl:     queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func (q *QueueOrder) SendMessageQueue(payment *entity.Payment) error {
	result, err := q.GetQueueURL()
	if err != nil {
		return err
	}

	queueURL := result.QueueUrl
	body, _ := json.Marshal(payment)

	err = q.SendMsg(queueURL, body)
	if err != nil {
		return err
	}
	return nil
}
