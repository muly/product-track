package main

import (
	"log"
	"os"

	"gopkg.in/mail.v2"
)

var emailClient *mail.Dialer

//function for initializing email 
func initEmailClient() {
	password := os.Getenv("GMAIL_PASSWORD")
	emailClient = mail.NewDialer("smtp.gmail.com", 587, "rohith.knaidu0125@gmail.com", password)
}

//function for sending email to the user according to the type of request
func sendEmail()(*mail.Message) {
	var t trackInput
	log.Println("creating mail")
	m := mail.NewMessage()
	m.SetHeader("From", "rohith.knaidu0125@gmail.com")
	m.SetHeader("To", "smulytestground@gmail.com")
	if err := emailClient.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return nil
	}
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		m.SetBody("text/plain", "product %s is available"+t.Url)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
	}
	return m 
}
