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
var secretManagerClient *secretmanager.Client
var password string

func initEmailClient() error {
	projectID := "149500152182" // project id in number format, not alpha string
	// Create the client.
	ctx := context.Background()
	var err error
	secretManagerClient, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("failed to setup email client: %v", err)
		return err
	}
	defer secretManagerClient.Close()

	secretID := "GMAIL_PASSWORD"
	if password = os.Getenv(secretID); password == "" {
		secretVersion, err := secretManagerClient.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
			Name: fmt.Sprintf("projects/%s/secrets/%s/versions/1", projectID, secretID),
		})
		if err != nil {
			log.Printf("failed to access secret version: %v", err)
			return err
		}
		password = string(secretVersion.Payload.Data)
	}
	emailClient = mail.NewDialer("smtp.gmail.com", 587, "rohith.knaidu0125@gmail.com", password)

	return nil
}

func sendEmail(t trackInput) error {
	log.Println("creating mail")
	m := mail.NewMessage()
	m.SetHeader("From", "rohith.knaidu0125@gmail.com")
	m.SetHeader("To", "smulytestground@gmail.com")
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		m.SetBody("text/plain", "product %s is available"+t.Url)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		m.SetBody("text/plain", "product %s is available with the minimum cost you needed"+t.Url)
	}
	if err := emailClient.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}
