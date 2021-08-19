package config

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

func TestQueuePublisher_PublishPayment(t *testing.T) {
	queue := "payment_queue" // default name for queue
	sess, _ := InitQueueSession()

	type args struct {
		payment model.Payments
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_publish_payment_to_sqs",
			args: args{
				payment: model.Payments{
					Name:          "Some test payment",
					TransactionId: "trx-00000-0000-0000",
					PaymentCode:   "payment-001",
					Amount:        100000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publish, err := NewQueuePublisher(sess, &queue)

			if err != nil {
				panic(err)
			}

			publish.Client.PurgeQueue(&sqs.PurgeQueueInput{
				QueueUrl: publish.QueueUrl,
			})

			if err := publish.PublishPayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("QueuePublisher.PublishPayment() error = %v, wantErr %v", err, tt.wantErr)
			}

			var timeout int64
			timeout = 10

			message, getErr := publish.GetPaymentMessage(&timeout)
			if getErr != nil {
				panic(getErr)
			}

			if len(message.Messages) == 0 {
				t.Error("GetPaymentMessage got 0 messages")
			}

			publish.Client.PurgeQueue(&sqs.PurgeQueueInput{
				QueueUrl: publish.QueueUrl,
			})
		})
	}
}
