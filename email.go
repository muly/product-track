package main

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"gopkg.in/mail.v2"
)

var emailClient *mail.Dialer

func initEmailClient() error {
	projectNumber := os.Getenv("PROJECT_NUMBER")

	secretID := "GMAIL_PASSWORD"
	password := os.Getenv(secretID)
	if password == "" {
		// Create the secret manager client.
		ctx := context.Background()
		secretManagerClient, err := secretmanager.NewClient(ctx)
		if err != nil {
			log.Printf("failed to setup email client: %v", err)
			return err
		}
		defer secretManagerClient.Close()

		// get secret
		secretVersion, err := secretManagerClient.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
			Name: fmt.Sprintf("projects/%s/secrets/%s/versions/1", projectNumber, secretID),
		})
		if err != nil {
			log.Printf("failed to access secret: %v", err)
			return err
		}
		password = string(secretVersion.Payload.Data)
	}
	emailClient = mail.NewDialer("smtp.gmail.com", 587, "rohith.knaidu0125@gmail.com", password)

	return nil
}

func sendTrackNotificationEmail(t trackInput) error {
	m, err := prepareTrackNotificationEmail(t)
	if err != nil {
		log.Println("Error preparing email:", err)
		return err
	}

	if err := emailClient.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

func prepareTrackNotificationEmail(t trackInput) (*mail.Message, error) {
	log.Println("creating mail")
	m := mail.NewMessage()
	m.SetHeader("From", "rohith.knaidu0125@gmail.com")
	m.SetHeader("To", emailid)
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		m.SetBody("text/plain", "product is available: "+t.Url)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product is available with the minimum cost you needed: "+t.Url)
	} else {
		return nil, fmt.Errorf("invalid request type %s", t.TypeOfRequest)
	}

	return m, nil
}
