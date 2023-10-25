package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"gopkg.in/mail.v2"
)

var emailClient *mail.Dialer

const systemEmailID = "rohith.knaidu0125@gmail.com"

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
	emailClient = mail.NewDialer("smtp.gmail.com", 587, systemEmailID, password)

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
	m.SetHeader("From", systemEmailID)
	m.SetHeader("To", t.EmailID)
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
		htmlBody := `<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; text-align: center; padding: 20px;">
			<div style="background-color: #ffffff; padding: 20px; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); max-width: 400px; margin: 0 auto;">
				<p style="font-weight: bold; color: #2d71ac;">Product is available:</p>
				<p style="color: #2d71ac;">Check out the product here <a href="PRODUCT_URL" style="color: #007bff; text-decoration: none; font-weight: bold;">Product's url</a></p>
			</div>
		</body>
		</html>`
		htmlBody = strings.Replace(htmlBody, "PRODUCT_URL", t.URL, -1)

		m.SetBody("text/html", htmlBody)
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
		htmlBody := `<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; text-align: center; padding: 20px;">
			<div style="background-color: #ffffff; padding: 20px; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); max-width: 400px; margin: 0 auto;">
				<p style="font-weight: bold; color: #2d71ac;">Product is available:</p>
				<p style="color: #2d71ac;">Check out the product here <a href="PRODUCT_URL" style="color: #007bff; text-decoration: none; font-weight: bold;">Product's url</a></p>
			</div>
		</body>
		</html>`
		htmlBody = strings.Replace(htmlBody, "PRODUCT_URL", t.URL, -1)
		m.SetBody("text/html", htmlBody)
	} else {
		return nil, fmt.Errorf("invalid request type %s", t.TypeOfRequest)
	}

	return m, nil
}
