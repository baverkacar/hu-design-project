package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
)

const (
	SmtpHost       = "smtp.gmail.com"
	SmtpPort       = "587"
	SenderMail     = "baverkacar@gmail.com"
	SenderPassword = "qizd cpul nruh ydvx"
)

func SendMail(email string, code string) (string, error) {
	token := SenderPassword
	var err error
	var msg []byte
	auth := smtp.PlainAuth("", SenderMail, token, SmtpHost)
	to := []string{email}
	if len(code) <= 0 {
		code, err = CreateVerificationCode()
		if err != nil {
			return "", err
		}
		msg = []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: Verification Code\r\n"+
			"\r\n"+
			"Verification code for activate your account: %s\r\n", email, code))
	} else {
		msg = []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: Reset Password Request\r\n"+
			"\r\n"+
			"New Password: %s\r\n", email, code))
	}

	err = smtp.SendMail(SmtpHost+":"+SmtpPort, auth, SenderMail, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return code, nil
}

func CreateVerificationCode() (string, error) {
	var code string
	for i := 0; i < 8; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += fmt.Sprintf("%d", num.Int64())
	}
	return code, nil
}
