package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/model"
	"github.com/yanoandri/yano-golang-training-beginner/services"
)

type PaymentCodeService struct {
	Repository services.Repository
}

type PaymentCodeRequest struct {
	Name        string `json:"name" validate:"required"`
	PaymentCode string `json:"payment_code" validate:"required"`
	Status      string `json:"status"`
}

func (service PaymentCodeService) CreatePaymentCode(c echo.Context) error {
	validate := validator.New()
	payment := new(PaymentCodeRequest)
	// bind the request body
	if err := c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// validate json required
	if err := validate.Struct(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// adding 50 years of expiration date
	dateExpired := time.Now().AddDate(0, 0, 90)
	//create database model
	paymentData := model.PaymentCodes{
		Name:           payment.Name,
		PaymentCode:    payment.PaymentCode,
		Status:         model.Active,
		ExpirationDate: fmt.Sprintf("%s", dateExpired.UTC()),
	}
	// save
	result, err := service.Repository.CreatePaymentCode(paymentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// return
	return c.JSON(http.StatusCreated, result)
}

func (service PaymentCodeService) GetPaymentCodeById(c echo.Context) error {
	id := c.Param("id")
	paymentCode, err := service.Repository.GetPaymentCodeById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, paymentCode)
}

func NewPaymentCodeController(e *echo.Echo, repo services.Repository) {
	controller := &PaymentCodeService{Repository: repo}
	e.POST("/payment-codes", controller.CreatePaymentCode)
	e.GET("/payment-codes/:id", controller.GetPaymentCodeById)
}
