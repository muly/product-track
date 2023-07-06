package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mail.v2"
)

func notify(t trackInput) error {
	m := mail.NewMessage()
	m.SetHeader("From", "smulytestground@gmail.com")
	m.SetHeader("To", "msrinivasareddy@gmail.com")
	if t.TypeOfRequest=="AVAILABILITY"{
	m.SetHeader("Subject", "Availability update Notification")
	m.SetBody("text/plain", "product %s is available"+t.Url)
	}else if t.TypeOfRequest=="PRICE"{
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
	}
	password := os.Getenv("GMAIL_PASSWORD")
	d := mail.NewDialer("smtp.gmail.com", 587, "smulytestground@gmail.com", password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	log.Printf("Email sent successfully %s request for %s", t.TypeOfRequest, t.Url)
	return nil
}
