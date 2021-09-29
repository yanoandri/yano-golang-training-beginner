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

func TestPaymentCodeService_CreatePaymentCode(t *testing.T) {
	e := echo.New()
	config.SetupDB()
	db := config.GetDBInstance()
	paymentCodeService := services.NewPaymentCodeService(db)

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
				Repository: paymentCodeService,
			},
			args: args{
				c: e,
			},
			payload: `{"payment_code":"PCODE-TEST-001","name":"Jon Snow", "status": "ACTIVE"}`,
			wantErr: false,
		},
		{
			name: "failed_create_payment_code_no_name_attributes",
			fields: fields{
				Repository: paymentCodeService,
			},
			args: args{
				c: e,
			},
			payload: `{"payment_code":"PCODE-TEST-001", "status": "ACTIVE"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/payment-codes", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.c.NewContext(req, rec)
			repo := tt.fields.Repository
			controller := PaymentCodeService{Repository: repo}
			if !tt.wantErr {
				if !assert.NoError(t, controller.CreatePaymentCode(c)) {
					t.Error("CreatePaymentCode integration test failed")
				} else {
					assert.Equal(t, http.StatusCreated, rec.Code)
				}
			} else {
				assert.Error(t, controller.CreatePaymentCode(c))
			}
		})
	}
	paymentCodeService.Database.Where("payment_code = ?", "PCODE-TEST-001").Delete(model.PaymentCodes{})
}

func TestPaymentCodeService_GetPaymentCodeById(t *testing.T) {
	e := echo.New()
	config.SetupDB()
	db := config.GetDBInstance()
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
		param   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success_get_payment_code_by_id",
			fields: fields{
				Repository: paymentCodeService,
			},
			args: args{
				c: e,
			},
			param:   paymentCodeModel.ID.String(),
			wantErr: false,
		},
		{
			name: "failed_not_found_get_payment_code_by_id",
			fields: fields{
				Repository: paymentCodeService,
			},
			args: args{
				c: e,
			},
			param:   "123",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/payment-codes/:id", nil)
			rec := httptest.NewRecorder()
			c := tt.args.c.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.param)
			repo := tt.fields.Repository
			controller := PaymentCodeService{Repository: repo}
			if !assert.NoError(t, controller.GetPaymentCodeById(c)) {
				t.Error("CreatePaymentCode integration test failed")
			} else {
				if !tt.wantErr {
					assert.Equal(t, http.StatusOK, rec.Code)
				} else {
					assert.Equal(t, http.StatusNotFound, rec.Code)
				}
			}
		})
	}
	paymentCodeService.Database.Where("id = ?", paymentCodeModel.ID.String()).Delete(model.PaymentCodes{})
}
