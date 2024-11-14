package middleware

import (
	"ezratoga/sendemail/config"
	"fmt"
	"strconv"
	gomail "gopkg.in/gomail.v2"
)

// sendEmail sends an email with the given subject, HTML content, and attachment file.
func SendEmail(subject, htmlContent, attachmentPath string) error {
	mailMiddleware := gomail.NewMessage()
	mailMiddleware.SetHeader("From", globalconfig.GetEnvVariable("SMTP_EMAIL")) // define smtp email in env file
	mailMiddleware.SetHeader("To", globalconfig.GetEnvVariable("EMAIL_RECEIVER")) // for testing you can set the receiver from env file
	mailMiddleware.SetBody("text/html", htmlContent)
	mailMiddleware.SetHeader("Subject", subject)

	if attachmentPath != "" { // handle if any attachment
		mailMiddleware.Attach(attachmentPath)
	}

	// convert string port from env to integer
	port, err := strconv.Atoi(globalconfig.GetEnvVariable("SMTP_PORT")) // Convert string to integer

    if err != nil {
        port = 587
    }

	d := gomail.NewDialer(globalconfig.GetEnvVariable("SMTP_HOST"), port, globalconfig.GetEnvVariable("SMTP_EMAIL"), globalconfig.GetEnvVariable("SMTP_PASSWORD"))

	// Send the email.
	if err := d.DialAndSend(mailMiddleware); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}