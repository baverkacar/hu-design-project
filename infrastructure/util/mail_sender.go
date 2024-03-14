package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"os"
)

const (
	SmtpHost   = "smtp.gmail.com"
	SmtpPort   = "587"
	SenderMail = "baverkacar@gmail.com"
)

func SendMail(email string) (string, error) {
	token := os.Getenv("SENDER_PASSWORD")

	code, err := createVerificationCode()
	if err != nil {
		return "", err
	}

	auth := smtp.PlainAuth("", SenderMail, token, SmtpHost)
	to := []string{email}
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Verification Code\r\n"+
		"\r\n"+
		"Verification code for activate your account: %s\r\n", email, code))
	err = smtp.SendMail(SmtpHost+":"+SmtpPort, auth, SenderMail, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return code, nil
}

func createVerificationCode() (string, error) {
	var code string
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += fmt.Sprintf("%d", num.Int64())
	}
	return code, nil
}
