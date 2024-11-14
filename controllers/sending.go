package controllers

import (
	"ezratoga/sendemail/models"
	"ezratoga/sendemail/helper"
    "net/http"
    "github.com/labstack/echo/v4"
	"os"
	"fmt"
)

// handleSendEmail is the handler for the POST /send-email endpoint.
func HandleSendEmail(c echo.Context) error {
	payload := new(models.EmailPayload)
	fmt.Print(*payload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Generate PDF from HTML attachment content.
	attachmentPath, err := middleware.GeneratePDF(payload.Attachment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate PDF"})
	}
	defer os.Remove(attachmentPath) // to remove file after sending email success

	// Send the email with HTML content and PDF attachment.
	if err := middleware.SendEmail(payload.Subject, payload.Content, attachmentPath); err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}
	fmt.Print(middleware.SendEmail)

	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
}