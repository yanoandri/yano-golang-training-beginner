package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/yanoandri/yano-golang-training-beginner/config"
	controller "github.com/yanoandri/yano-golang-training-beginner/controller"
	service "github.com/yanoandri/yano-golang-training-beginner/services"
	"gorm.io/gorm"
)

type Healthy struct {
	Status string `json:"status"`
}

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	config.SetupDB()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "web":
			sess, err := config.InitQueueSession()
			if err != nil {
				// do not
				fmt.Sprintln(err)
			}
			web(config.GetDBInstance(), sess)
		case "cron":
			cron(config.GetDBInstance())
		default:
			fmt.Println("Command unknown...")
		}
	}

}

func web(db *gorm.DB, session *session.Session) {
	e := echo.New()
	queue := "payment_queue"
	result, _ := config.GetQueueURL(session, &queue)
	fmt.Printf("queue url is: %v", result.QueueUrl)
	config.SendMsg(session, result.QueueUrl)

	e.GET("/hello-world", helloWorld)
	e.GET("/health", healthy)
	paymentCodeService := service.NewPaymentCodeService(db)
	controller.NewPaymentCodeController(e, paymentCodeService)
	inquiryService := service.NewInquiryService(db)
	controller.NewInquiryController(e, inquiryService)
	paymentService := service.NewPaymentService(db)
	controller.NewPaymentController(e, paymentService)
	e.Logger.Fatal(e.Start(":1323"))
}

func cron(db *gorm.DB) {
	fmt.Println("...Cron Start...")
	paymentCodeService := service.NewPaymentCodeService(db)
	result := paymentCodeService.ExpirePaymentCode()
	fmt.Printf("Number of Row Affected : %d\n", result)
	fmt.Println("...Cron Ended...")
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
