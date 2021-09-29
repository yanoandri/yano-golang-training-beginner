package controller

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/model"
	"github.com/yanoandri/yano-golang-training-beginner/services"
)

type PaymentService struct {
	Repository services.Repository
}

type PaymentRequest struct {
	TransactionId string `json:"transaction_id" validate:"required"`
	PaymentCode   string `json:"payment_code" validate:"required"`
	Amount        int64  `json:"amount" validate:"required"`
	Name          string `json:"name" validate:"required"`
}

func (service PaymentService) CreatePayment(c echo.Context) error {
	validate := validator.New()
	payment := new(PaymentRequest)
	// bind the request body
	if err := c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// validate json required
	if err := validate.Struct(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check if payment code is exists and active
	_, paymentCodeErr := service.Repository.GetActivePaymentCodeByCode(payment.PaymentCode)
	if paymentCodeErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, paymentCodeErr)
	}

	// check if transaaction id is exists
	_, inquiryErr := service.Repository.GetInquiryByTransactionId(payment.TransactionId)
	if inquiryErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, inquiryErr)
	}

	//create database model
	paymentData := model.Payments{
		TransactionId: payment.TransactionId,
		PaymentCode:   payment.PaymentCode,
		Name:          payment.Name,
		Amount:        payment.Amount,
	}
	// save
	result, err := service.Repository.CreatePayment(paymentData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	type SuccessPayment struct {
		ID            string `json:"id"`
		TransactionId string `json:"transaction_id"`
		Name          string `json:"name"`
		Status        string `json:"status"`
	}

	successResponse := &SuccessPayment{
		ID:            result.ID.String(),
		TransactionId: result.TransactionId,
		Name:          result.Name,
		Status:        "SUCCESS",
	}
	// return
	return c.JSON(http.StatusOK, successResponse)
}

func NewPaymentController(e *echo.Echo, repo services.Repository) {
	controller := &PaymentService{Repository: repo}
	e.POST("/payment", controller.CreatePayment)
}
