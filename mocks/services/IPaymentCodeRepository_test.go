// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	model "github.com/yanoandri/yano-golang-training-beginner/model"
)

func TestIPaymentCodeRepository_Create(t *testing.T) {
	dateExpired := fmt.Sprintf("%s", time.Now().AddDate(50, 0, 0).UTC())
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		payment model.PaymentCodes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.PaymentCodes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Success create payment code",
			fields: fields{},
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
			name:   "Failed create payment code",
			fields: fields{},
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
			_m := &IPaymentCodeRepository{
				Mock: tt.fields.Mock,
			}

			if !tt.wantErr {
				_m.On("Create", tt.args.payment).Return(tt.want, nil)
			} else {
				_m.On("Create", tt.args.payment).Return(tt.want, errors.New("errors"))
			}

			got, err := _m.Create(tt.args.payment)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPaymentCodeRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPaymentCodeRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPaymentCodeRepository_GetPaymentById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.PaymentCodes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Success get payment code by id",
			fields: fields{},
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
			name:   "Failed get payment code by id",
			fields: fields{},
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    model.PaymentCodes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &IPaymentCodeRepository{
				Mock: tt.fields.Mock,
			}

			if !tt.wantErr {
				_m.On("GetPaymentById", tt.args.id).Return(tt.want, nil)
			} else {
				_m.On("GetPaymentById", tt.args.id).Return(tt.want, errors.New("errors"))
			}

			got, err := _m.GetPaymentById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPaymentCodeRepository.GetPaymentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPaymentCodeRepository.GetPaymentById() = %v, want %v", got, tt.want)
			}
		})
	}
}