package main

import (
	"fmt"
	"net/smtp"
)

func sendEmail(poem string) {
	from := "rumi@poems.com"
	to := []string{"andres.crucetta@hey.com"}
	smtpHost := "localhost"
	smtpPort := "1025"

	subject := "Daily Rumi Poem"

	body := fmt.Sprintf("Here's your daily Rumi poem:\n\n%s", poem)
	message := []byte(fmt.Sprintf("Subject: %s\n\n%s", subject, body))

	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
}
