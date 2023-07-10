package main

import (
	"log"

	"gopkg.in/mail.v2"
)

var emailClient *mail.Dialer

// function for initializing email
func initEmailClient() {
	password :=string(payloadData)
	emailClient = mail.NewDialer("smtp.gmail.com", 587, "smulytestground@gmail.com", password)
}

// function for sending email to the user according to the type of request
func sendEmail(t trackInput) error {
	log.Println("creating mail")
	m := mail.NewMessage()
	m.SetHeader("From", "smulytestground@gmail.com")
	m.SetHeader("To", "rohith.knaidu0125@gmail.com")
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		m.SetBody("text/plain", "product %s is available"+t.Url)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
	}
	if err := emailClient.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
	}

	return nil
}
