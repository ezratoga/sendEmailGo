package models

type EmailPayload struct {
	Subject    string `json:"subject" validate:"required`
	Content    string `json:"content" validate:"required`
	Attachment string `json:"attachment" validate:"required`
}