package config

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

type IQueuePublisher interface {
	PublishPayment(payment model.Payments) error
	GetPaymentMessage(timeout *int64) (*sqs.ReceiveMessageOutput, error)
}

type QueuePublisher struct {
	Client   *sqs.SQS
	QueueUrl *string
}

func NewQueuePublisher(sess *session.Session, queue *string) (QueuePublisher, error) {
	client := sqs.New(sess)

	res, err := client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})

	if err != nil {
		return QueuePublisher{}, err
	}

	return QueuePublisher{Client: client, QueueUrl: res.QueueUrl}, nil
}

func (queue QueuePublisher) PublishPayment(payment model.Payments) error {
	message, _ := json.Marshal(payment)

	_, err := queue.Client.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(message)),
		QueueUrl:    queue.QueueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}

func InitQueueSession() (*session.Session, error) {
	// Create a session that gets credential values
	region := "ap-southeast-2"
	endpoint := "http://localhost:4566/000000000000/payment_queue"

	sess, err := session.NewSession(&aws.Config{
		Region:   &region,
		Endpoint: &endpoint,
	})

	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (queue QueuePublisher) GetPaymentMessage(timeout *int64) (*sqs.ReceiveMessageOutput, error) {
	// snippet-start:[sqs.go.receive_messages.call]
	msgResult, err := queue.Client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		QueueUrl:            queue.QueueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   timeout,
	})
	// snippet-end:[sqs.go.receive_messages.call]
	if err != nil {
		return nil, err
	}

	return msgResult, nil
}
