package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Healthy struct {
	Status string `json:"status"`
}

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/hello-world", helloWorld)
	e.GET("/health", healthy)
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
