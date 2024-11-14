package middleware

import (
	"bytes"
	"os/exec"
	"time"
	"fmt"
)

func GeneratePDF(HtmlContent string) (string, error) {
	filename := fmt.Sprintf("attachment_%d.pdf", time.Now().Unix())
	cmd := exec.Command("wkhtmltopdf", "-", filename)
	cmd.Stdin = bytes.NewBufferString(HtmlContent)

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}
	return filename, nil
}