package utils

import (
	"fmt"
)

func SendResetPasswordEmail(email, token string) error {
	subject := "Reset Password Email"
	body := fmt.Sprintf("Please reset your password by clicking the following link: http://localhost:8080/reset-password?token=%s", token)
	return sendEmail(email, subject, body)
}
