package middleware

import (
	"bytes"
	"os/exec"
	"time"
	"fmt"
)

// helper function for generate pdf
func GeneratePDF(HtmlContent string) (string, error) {
	filename := fmt.Sprintf("attachment_%d.pdf", time.Now().Unix()) // define filename to send int an email
	cmd := exec.Command("wkhtmltopdf", "-", filename) // run command wkhtmltopdf in the server to generate pdf from html
	cmd.Stdin = bytes.NewBufferString(HtmlContent) // add buffer string input from html that attached in payload

	err := cmd.Run() // generate pdf
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}
	return filename, nil
}