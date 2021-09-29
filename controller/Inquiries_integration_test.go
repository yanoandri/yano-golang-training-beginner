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

func TestInquiryService_CreateInquiry(t *testing.T) {
	e := echo.New()
	config.SetupDB()
	db := config.GetDBInstance()
	inquiryService := services.NewInquiryService(db)
	paymentCodeService := services.NewPaymentCodeService(db)

	paymentCodeModel := &model.PaymentCodes{
		Name:        "Payment code name",
		PaymentCode: "PCODE-TEST-001",
		Status:      "ACTIVE",
	}

	paymentCodeService.Database.Create(&paymentCodeModel)

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
			name: "success_create_payment_code",
			fields: fields{
				Repository: inquiryService,
			},
			args: args{
				c: e,
			},
			payload: `{"transaction_id":"TRX-TEST-ID-001","payment_code":"PCODE-TEST-001"}`,
			wantErr: false,
		},
		{
			name: "failed_create_payment_code_no_transaction_id_attribute",
			fields: fields{
				Repository: inquiryService,
			},
			args: args{
				c: e,
			},
			payload: `{"payment_code":"PCODE-TEST-001"}`,
			wantErr: true,
		},
		{
			name: "failed_create_payment_code_no_payment_code_active_found",
			fields: fields{
				Repository: inquiryService,
			},
			args: args{
				c: e,
			},
			payload: `{"payment_code":"PCODE-TEST-002"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/inquiry", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.c.NewContext(req, rec)
			repo := tt.fields.Repository
			controller := InquiryService{Repository: repo}
			if !tt.wantErr {
				if !assert.NoError(t, controller.CreateInquiry(c)) {
					t.Error("CreateInquiry integration test failed")
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
				}
			} else {
				assert.Error(t, controller.CreateInquiry(c))
			}
		})
	}
	paymentCodeService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.PaymentCodes{})
	inquiryService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.Inquiries{})
}
