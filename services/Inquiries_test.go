package services

import (
	"errors"
	"reflect"
	"testing"

	mocks "github.com/yanoandri/yano-golang-training-beginner/mocks/services"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

func TestRepository_CreateInquiry(t *testing.T) {
	type args struct {
		inquiry model.Inquiries
	}
	tests := []struct {
		name    string
		args    args
		want    model.Inquiries
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_create_inquiry",
			args: args{
				model.Inquiries{
					TransactionId: "TRX-123",
					PaymentCode:   "XXX-123",
				},
			},
			want: model.Inquiries{
				TransactionId: "TRX-123",
				PaymentCode:   "XXX-123",
			},
			wantErr: false,
		},
		{
			name: "failed_create_inquiry",
			args: args{
				model.Inquiries{
					TransactionId: "TRX-123",
					PaymentCode:   "XXX-123",
				},
			},
			want:    model.Inquiries{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IInquiriesService{}
			if !tt.wantErr {
				conn.On("CreateInquiry", tt.args.inquiry).Return(tt.want, nil)
			} else {
				conn.On("CreateInquiry", tt.args.inquiry).Return(model.Inquiries{}, errors.New("inquiry error"))
			}

			got, err := conn.CreateInquiry(tt.args.inquiry)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateInquiry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreateInquiry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetInquiryByTransactionId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Inquiries
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_get_inquiry_by_transaction_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want: model.Inquiries{
				TransactionId: "TRX-123",
				PaymentCode:   "XXX-123",
			},
			wantErr: false,
		},
		{
			name: "failed_get_inquiry_by_transaction_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    model.Inquiries{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.IInquiriesService{}

			if !tt.wantErr {
				conn.On("GetInquiryByTransactionId", tt.args.id).Return(tt.want, nil)
			} else {
				conn.On("GetInquiryByTransactionId", tt.args.id).Return(model.Inquiries{}, errors.New("inquiry error"))
			}

			got, err := conn.GetInquiryByTransactionId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetInquiryByTransactionId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetInquiryByTransactionId() = %v, want %v", got, tt.want)
			}
		})
	}
}
