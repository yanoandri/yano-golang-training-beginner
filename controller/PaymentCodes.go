package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

type IPaymentCodeService interface {
	CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error)
	Get(id string) (model.PaymentCodes, error)
}

type PaymentCodeService struct {
	PaymentCodeService IPaymentCodeService
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
	dateExpired := time.Now().AddDate(50, 0, 0)
	//create database model
	paymentData := model.PaymentCodes{
		Name:           payment.Name,
		PaymentCode:    payment.PaymentCode,
		Status:         model.Active,
		ExpirationDate: fmt.Sprintf("%s", dateExpired.UTC()),
	}
	// save
	result, err := service.PaymentCodeService.CreatePaymentCode(paymentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	// return
	return c.JSON(http.StatusCreated, result)
}

func (service PaymentCodeService) GetPaymentCodeById(c echo.Context) error {
	id := c.Param("id")
	paymentCode, err := service.PaymentCodeService.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, paymentCode)
}

func NewPaymentCodeController(e *echo.Echo, service IPaymentCodeService) {
	controller := &PaymentCodeService{PaymentCodeService: service}
	e.POST("/payment-codes", controller.CreatePaymentCode)
	e.GET("/payment-codes/:id", controller.GetPaymentCodeById)
}
