package main

import (
	// "bytes"
	// "fmt"
	// "net/http"
	// "os"
	// "os/exec"
	// "time"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ezratoga/sendemail/controllers"
	// gomail "gopkg.in/gomail.v2"
)

// CustomValidator for validation payload inputted are matching with the struct or not
type CustomValidator struct {
	validator *validator.Validate
}

// Validate makes CustomValidator satisfy the echo.Validator interface.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// initiate the server
	echoServer := echo.New()
	echoServer.Use(middleware.CORS()) // handle API Cors
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	// Register the custom validator
	echoServer.Validator = &CustomValidator{validator: validator.New()}

	// Routes for send email
	echoServer.POST("/send-email", controllers.HandleSendEmail)

	// Start server
	echoServer.Logger.Fatal(echoServer.Start(":8899"))
}
