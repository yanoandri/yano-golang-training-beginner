package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/config"
	"github.com/yanoandri/yano-golang-training-beginner/model"
)

type PaymentCodeRequest struct {
	Name        string `json:"name" validate:"required"`
	PaymentCode string `json:"payment_code" validate:"required"`
	Status      string `json:"status"`
}

func CreatePaymentCode(c echo.Context) error {
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
	config.GetDBInstance().Create(&paymentData)
	// return
	return c.JSON(http.StatusCreated, paymentData)
}

func UpdatePaymentCode(c echo.Context) error {
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
	//get payment code data
	id := c.Param("id")
	paymentCodes, err := getPaymentDataById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	// updating data
	config.GetDBInstance().Model(&paymentCodes).Updates(model.PaymentCodes{
		PaymentCode: payment.PaymentCode,
		Name:        payment.Name,
		Status:      payment.Status,
	})
	// return
	return c.JSON(http.StatusOK, paymentCodes)
}

func DeletePaymentCodes(c echo.Context) error {
	//get payment code data
	id := c.Param("id")
	paymentCodes, err := getPaymentDataById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	// updating data
	config.GetDBInstance().Delete(&paymentCodes)
	// return
	return c.JSON(http.StatusOK, paymentCodes)
}

func GetPaymentCodeById(c echo.Context) error {
	id := c.Param("id")
	paymentCode, err := getPaymentDataById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, paymentCode)
}

func getPaymentDataById(id string) (model.PaymentCodes, error) {
	var paymentData model.PaymentCodes
	result := config.GetDBInstance().First(&paymentData, "id = ?", id)
	if result.RowsAffected == 0 {
		return model.PaymentCodes{}, errors.New("Data not found")
	}
	return paymentData, nil
}
