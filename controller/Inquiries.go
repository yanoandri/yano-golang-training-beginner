package controller

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/model"
	"github.com/yanoandri/yano-golang-training-beginner/services"
)

type InquiryService struct {
	Repository services.Repository
}

type InquiryRequest struct {
	TransactionId string `json:"transaction_id" validate:"required"`
	PaymentCode   string `json:"payment_code" validate:"required"`
}

func (service InquiryService) CreateInquiry(c echo.Context) error {
	validate := validator.New()
	inquiry := new(InquiryRequest)
	// bind the request body
	if err := c.Bind(inquiry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// validate json required
	if err := validate.Struct(inquiry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// check if payment code is exists and active
	_, err := service.Repository.GetActivePaymentCodeByCode(inquiry.PaymentCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	//create database model
	inquiryData := model.Inquiries{
		TransactionId: inquiry.TransactionId,
		PaymentCode:   inquiry.PaymentCode,
	}
	// save
	result, err := service.Repository.CreateInquiry(inquiryData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	//get join table
	inquiryCodes, _ := service.Repository.GetInquiryByTransactionId(result.TransactionId)
	// return
	return c.JSON(http.StatusOK, inquiryCodes)
}

func NewInquiryController(e *echo.Echo, repo services.Repository) {
	controller := &InquiryService{Repository: repo}
	e.POST("/inquiry", controller.CreateInquiry)
}
