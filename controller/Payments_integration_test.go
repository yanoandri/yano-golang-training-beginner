package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yanoandri/yano-golang-training-beginner/config"
	"github.com/yanoandri/yano-golang-training-beginner/model"
	"github.com/yanoandri/yano-golang-training-beginner/services"
)

func TestPaymentService_CreatePayment(t *testing.T) {
	e := echo.New()
	config.SetupDB()
	db := config.GetDBInstance()
	paymentService := services.NewPaymentService(db)
	inquiryService := services.NewInquiryService(db)
	paymentCodeService := services.NewPaymentCodeService(db)

	paymentCodeModel := &model.PaymentCodes{
		Name:        "Payment code name",
		PaymentCode: "PCODE-TEST-001",
		Status:      "ACTIVE",
	}

	inquiryModel := &model.Inquiries{
		TransactionId: "TRX-TEST-001",
		PaymentCode:   paymentCodeModel.PaymentCode,
	}

	paymentCodeService.Database.Create(&paymentCodeModel)
	inquiryService.Database.Create(&inquiryModel)

	type fields struct {
		Repository services.Repository
	}
	type args struct {
		c *echo.Echo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		payload string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_create_payment",
			fields: fields{
				Repository: paymentService,
			},
			args: args{
				c: e,
			},
			payload: `{"transaction_id":"TRX-TEST-001","payment_code":"PCODE-TEST-001", "name": "Test Payment", "amount": 50000}`,
			wantErr: false,
		},
		{
			name: "failed_create_payment_code_no_amount_attribute",
			fields: fields{
				Repository: paymentService,
			},
			args: args{
				c: e,
			},
			payload: `{"transaction_id":"TRX-TEST-001","payment_code":"PCODE-TEST-001", "name": "Test Payment"}`,
			wantErr: true,
		},
		{
			name: "failed_create_payment_code_no_payment_code_active_found",
			fields: fields{
				Repository: paymentService,
			},
			args: args{
				c: e,
			},
			payload: `{"transaction_id":"TRX-TEST-001", "payment_code":"PCODE-TEST-002", "name": "Test Payment", "amount": 50000}`,
			wantErr: true,
		},
		{
			name: "failed_create_payment_code_no_transaction_id_found",
			fields: fields{
				Repository: paymentService,
			},
			args: args{
				c: e,
			},
			payload: `{"transaction_id":"TRX-TEST-ID-001", "payment_code":"PCODE-TEST-002", "name": "Test Payment", "amount": 50000}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.c.NewContext(req, rec)
			repo := tt.fields.Repository
			controller := PaymentService{Repository: repo}
			if !tt.wantErr {
				if !assert.NoError(t, controller.CreatePayment(c)) {
					t.Error("CreatePayment integration test failed")
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
				}
			} else {
				assert.Error(t, controller.CreatePayment(c))
			}
		})
	}
	paymentCodeService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.PaymentCodes{})
	inquiryService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.Inquiries{})
	paymentService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.Payments{})
}
