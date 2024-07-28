package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

var (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	smtpUser     = os.Getenv("SMTP_USER")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
)

func SendVerificationEmail(email, token string) error {
	from := smtpUser
	to := email
	subject := "Email Verification"
	//plainTextContent := fmt.Sprintf("Please verify your email using this token: %s", token)
	htmlContent := fmt.Sprintf("<p>Please verify your email using this token: <button><strong>%s</strong></button></p>", token)

	// Create the email message
	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s", from, to, subject, htmlContent))

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
