package services

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	mocks "github.com/yanoandri/yano-golang-training-beginner/mocks/services"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

func TestConfiguration_CreatePaymentCode(t *testing.T) {
	dateExpired := fmt.Sprintf("%s", time.Now().AddDate(50, 0, 0).UTC())
	type args struct {
		payment model.PaymentCodes
	}
	tests := []struct {
		name    string
		args    args
		want    model.PaymentCodes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_create_payment_code",
			args: args{
				model.PaymentCodes{
					Name:           "lechsa",
					PaymentCode:    "XXX-123",
					Status:         model.Active,
					ExpirationDate: dateExpired,
				},
			},
			want: model.PaymentCodes{
				Name:           "lechsa",
				PaymentCode:    "XXX-123",
				Status:         model.Active,
				ExpirationDate: dateExpired,
			},
			wantErr: false,
		},
		{
			name: "failed_create_payment_code",
			args: args{
				model.PaymentCodes{
					Name:           "lechsa",
					PaymentCode:    "XXX-123",
					Status:         model.Active,
					ExpirationDate: dateExpired,
				},
			},
			want:    model.PaymentCodes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := mocks.IPaymentCodeService{}

			if !tt.wantErr {
				conn.On("CreatePaymentCode", tt.args.payment).Return(tt.want, nil)
			} else {
				conn.On("CreatePaymentCode", tt.args.payment).Return(model.PaymentCodes{}, errors.New("payment error"))
			}

			got, err := conn.CreatePaymentCode(tt.args.payment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Configuration.CreatePaymentCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Configuration.CreatePaymentCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfiguration_GetPaymentCodeById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.PaymentCodes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_get_payment_code_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want: model.PaymentCodes{
				Name:           "lechsa",
				PaymentCode:    "XXX-123",
				Status:         model.Active,
				ExpirationDate: fmt.Sprintf("%s", time.Now().AddDate(50, 0, 0).UTC()),
			},
			wantErr: false,
		},
		{
			name: "failed_get_payment_code_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    model.PaymentCodes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IPaymentCodeService{}

			if !tt.wantErr {
				conn.On("GetPaymentCodeById", tt.args.id).Return(tt.want, nil)
			} else {
				conn.On("GetPaymentCodeById", tt.args.id).Return(model.PaymentCodes{}, errors.New("payment error"))
			}

			got, err := conn.GetPaymentCodeById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Configuration.GetPaymentCodeById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Configuration.GetPaymentCodeById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_ExpirePaymentCode(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "should_expire_one_payment_code",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := mocks.IPaymentCodeService{}
			conn.On("ExpirePaymentCode").Return(tt.want)
			if got := conn.ExpirePaymentCode(); got != tt.want {
				t.Errorf("Repository.ExpirePaymentCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetActivePaymentCodeByCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    model.PaymentCodes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_get_active_payment_code_by_code",
			args: args{
				code: "0000-0000-0000-000",
			},
			want: model.PaymentCodes{
				Name:           "lechsa",
				PaymentCode:    "XXX-123",
				Status:         model.Active,
				ExpirationDate: fmt.Sprintf("%s", time.Now().AddDate(50, 0, 0).UTC()),
			},
			wantErr: false,
		},
		{
			name: "failed_get_active_payment_code_by_code",
			args: args{
				code: "0000-0000-0000-000",
			},
			want:    model.PaymentCodes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IPaymentCodeService{}

			if !tt.wantErr {
				conn.On("GetActivePaymentCodeByCode", tt.args.code).Return(tt.want, nil)
			} else {
				conn.On("GetActivePaymentCodeByCode", tt.args.code).Return(model.PaymentCodes{}, errors.New("payment error"))
			}

			got, err := conn.GetActivePaymentCodeByCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetActivePaymentCodeByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetActivePaymentCodeByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
