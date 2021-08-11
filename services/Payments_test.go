package services

import (
	"errors"
	"reflect"
	"testing"

	mocks "github.com/yanoandri/yano-golang-training-beginner/mocks/services"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

func TestRepository_CreatePayment(t *testing.T) {
	type args struct {
		payment model.Payments
	}
	tests := []struct {
		name    string
		args    args
		want    model.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_create_payment",
			args: args{
				model.Payments{
					TransactionId: "TRX-123",
					PaymentCode:   "XXX-123",
					Amount:        1000,
					Name:          "Some name",
				},
			},
			want: model.Payments{
				TransactionId: "TRX-123",
				PaymentCode:   "XXX-123",
				Amount:        1000,
				Name:          "Some name",
			},
			wantErr: false,
		},
		{
			name: "failed_create_payment",
			args: args{
				model.Payments{
					TransactionId: "TRX-123",
					PaymentCode:   "XXX-123",
					Amount:        1000,
					Name:          "Some name",
				},
			},
			want:    model.Payments{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IPaymentsService{}

			if !tt.wantErr {
				conn.On("CreatePayment", tt.args.payment).Return(tt.want, nil)
			} else {
				conn.On("CreatePayment", tt.args.payment).Return(model.Payments{}, errors.New("payment error"))
			}

			got, err := conn.CreatePayment(tt.args.payment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetPaymentById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_get_payment_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want: model.Payments{
				TransactionId: "TRX-123",
				PaymentCode:   "XXX-123",
				Amount:        1000,
				Name:          "Some name",
			},
			wantErr: false,
		},
		{
			name: "failed_get_payment_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    model.Payments{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IPaymentsService{}

			if !tt.wantErr {
				conn.On("GetPaymentById", tt.args.id).Return(tt.want, nil)
			} else {
				conn.On("GetPaymentById", tt.args.id).Return(model.Payments{}, errors.New("payment error"))
			}

			got, err := conn.GetPaymentById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetPaymentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetPaymentById() = %v, want %v", got, tt.want)
			}
		})
	}
}
