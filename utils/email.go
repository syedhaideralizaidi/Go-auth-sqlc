package utils

import (
	"fmt"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

var sendgridAPIKey = os.Getenv("SENDGRID_API_KEY")

func SendVerificationEmail(email, token string) error {
	//from := mail.NewEmail("AssetCurve", "8info@assetcurve.io")
	from := mail.NewEmail("AssetCurve", "test@me.io")
	subject := "Email Verification"
	to := mail.NewEmail("", email)
	plainTextContent := fmt.Sprintf("Please verify your email using this token: %s", token)
	htmlContent := fmt.Sprintf("<p>Please verify your email using this token: <strong>%s</strong></p>", token)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	_, err := client.Send(message)
	return err
}
