package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/config"
	_paymentCodeController "github.com/yanoandri/yano-golang-training-beginner/controller"
	"github.com/yanoandri/yano-golang-training-beginner/repository"
	_paymentCodeService "github.com/yanoandri/yano-golang-training-beginner/services"
)

type Healthy struct {
	Status string `json:"status"`
}

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	config.NewDB()
	e.GET("/hello-world", helloWorld)
	e.GET("/health", healthy)

	paymentCodeRepository := repository.NewPaymentCodeRepository(config.GetDBInstance())
	paymentCodeService := _paymentCodeService.NewPaymentCodeService(*paymentCodeRepository)
	_paymentCodeController.NewPaymentCodeController(e, paymentCodeService)
	// e.POST("/payment-codes", controller.CreatePaymentCode)
	// e.PATCH("/payment-codes/:id", controller.UpdatePaymentCode)
	// e.DELETE("/payment-codes/:id", controller.DeletePaymentCodes)
	// e.GET("/payment-codes/:id", controller.GetPaymentCodeById)
	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorld(c echo.Context) error {
	message := &HelloWorld{
		Message: "hello world",
	}
	return c.JSON(http.StatusOK, message)
}

func healthy(c echo.Context) error {
	status := &Healthy{
		Status: "healthy",
	}
	return c.JSON(http.StatusOK, status)
}
