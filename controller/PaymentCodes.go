package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/model/request"
)

func CreatePaymentCode(c echo.Context) error {
	payment := new(request.PaymentCodes)

	if err := c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
