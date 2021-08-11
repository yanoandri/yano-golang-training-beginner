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
	TransactionId string `json:"name" validate:"required"`
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check if payment code is exists and active
	_, err := service.Repository.GetActivePaymentCodeByCode(inquiry.PaymentCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	//create database model
	inquiryData := model.Inquiries{
		TransactionId: inquiry.TransactionId,
		PaymentCode:   inquiry.PaymentCode,
	}
	// save
	result, err := service.Repository.CreateInquiry(inquiryData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	// return
	return c.JSON(http.StatusCreated, result)
}

func NewInquiryController(e *echo.Echo, repo services.Repository) {
	controller := &InquiryService{Repository: repo}
	e.POST("/inquiry", controller.CreateInquiry)
}
