package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mail.v2"
)

func initemail() ( *mail.Message,error) {
	m := mail.NewMessage()
	m.SetHeader("From", "smulytestground@gmail.com")
	m.SetHeader("To", "msrinivasareddy@gmail.com")
	
	password := os.Getenv("GMAIL_PASSWORD")
	d := mail.NewDialer("smtp.gmail.com", 587, "smulytestground@gmail.com", password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return nil,err
	}

	log.Printf("Email sent successfully")
	return m,nil
}

func requesttype( *mail.Message){
	var t trackInput
	m,err:=initemail()
	if err!=nil{
		log.Println("error during ",err)
	}else{
		if t.TypeOfRequest==requestTypeAvailability{
			m.SetHeader("Subject", "Availability update Notification")
			m.SetBody("text/plain", "product %s is available"+t.Url)
		}else if t.TypeOfRequest==requestTypePrice{
			m.SetHeader("Subject", "price update Notification")
			m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
		}
	} 
}

