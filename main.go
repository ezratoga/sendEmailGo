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

// EmailPayload represents the JSON structure of the incoming payload.
// type EmailPayload struct {
// 	Subject    string `json:"subject" validate:"required`
// 	Content    string `json:"content" validate:"required`
// 	Attachment string `json:"attachment" validate:"required`
// }

// generatePDF converts HTML content to a PDF file using wkhtmltopdf.
// func generatePDF(htmlContent string) (string, error) {
// 	filename := fmt.Sprintf("attachment_%d.pdf", time.Now().Unix())
// 	cmd := exec.Command("wkhtmltopdf", "-", filename)
// 	cmd.Stdin = bytes.NewBufferString(htmlContent)

// 	err := cmd.Run()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate PDF: %w", err)
// 	}
// 	return filename, nil
// }

// sendEmail sends an email with the given subject, HTML content, and attachment file.
// func sendEmail(subject, htmlContent, attachmentPath string) error {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "MS_fEjTvz@trial-jy7zpl9kdwo45vx6.mlsender.net")
// 	m.SetHeader("To", "ezra.tphtbrt@gmail.com")
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", htmlContent)
// 	m.Attach(attachmentPath)

// 	d := gomail.NewDialer("smtp.mailersend.net", 587, "MS_fEjTvz@trial-jy7zpl9kdwo45vx6.mlsender.net", "oGKPgYt4SS6tpMET")

// 	// Send the email.
// 	if err := d.DialAndSend(m); err != nil {
// 		return fmt.Errorf("failed to send email: %w", err)
// 	}
// 	return nil
// }

// handleSendEmail is the handler for the POST /send-email endpoint.
// func handleSendEmail(c echo.Context) error {
// 	payload := new(EmailPayload)
// 	fmt.Print(payload)

// 	if err := c.Bind(payload); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
// 	}
// 	if err := c.Validate(payload); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	}

// 	// Generate PDF from HTML attachment content.
// 	attachmentPath, err := generatePDF(payload.Attachment)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate PDF"})
// 	}
// 	defer os.Remove(attachmentPath)

// 	// Send the email with HTML content and PDF attachment.
// 	if err := sendEmail(payload.Subject, payload.Content, attachmentPath); err != nil {
// 		fmt.Print(err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
// 	}
// 	fmt.Print(sendEmail)

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
// }

// CustomValidator wraps the validator instance from go-playground/validator.
type CustomValidator struct {
	validator *validator.Validate
}

// Validate makes CustomValidator satisfy the echo.Validator interface.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	echoServer := echo.New()
	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	// Register the custom validator
	echoServer.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	echoServer.POST("/send-email", controllers.HandleSendEmail)

	// Start server
	echoServer.Logger.Fatal(echoServer.Start(":8899"))
}
