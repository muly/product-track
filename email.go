package main

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

var emailclient *mail.Dialer

func initEmailClient() {

	password := os.Getenv("GMAIL_PASSWORD")
	emailclient = mail.NewDialer("smtp.gmail.com", 587, "smulytestground@gmail.com", password)

}

func requesttype(*mail.Message) {
	var t trackInput
	m := mail.NewMessage()
	m.SetHeader("From", "smulytestground@gmail.com")
	m.SetHeader("To", "rohith.knaidu0125@gmail.com")
	// if err!=nil{
	// 	log.Println("error during ",err)
	// }
	if err := emailclient.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		m.SetBody("text/plain", "product %s is available"+t.Url)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
	}
}
