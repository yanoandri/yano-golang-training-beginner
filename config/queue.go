package config

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

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

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendMsg(sess *session.Session, queueURL *string) error {
	// Create an SQS service client
	svc := sqs.New(sess)

	type exampleMessage struct {
		Name string
	}

	message := exampleMessage{"John Doe"}
	messageBody, _ := json.Marshal(message)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}
