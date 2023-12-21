package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	scrape "github.com/muly/product-scrape"
	"gopkg.in/mail.v2"
)

const (
	envGmailID       = "GMAIL_ID"
	envGmailPassword = "GMAIL_PASSWORD"
	envProjectNumber = "PROJECT_NUMBER"
)

const notificationEmailBody = `<html>
<body style="font-family: Arial, Helvetica, sans-serif; text-align: center;">
  <div style="background-color: #ffffff; padding: 10px 20px; border-radius: 10px; max-width: 400px; margin: 0 auto;">
  <img style="max-width: 100%; margin: 3px 0 0; border-radius: 8px; width: 85px; height: 85px;" src="cid:logo.png" alt="Logo">
	<h2 style="font-weight: bolder; color: #2d71ac; font-size: 20px;">Product is available</h2>
	<img style="max-width: 100%; margin: 3px 0; border-radius: 8px;" src="product-image" alt="Product Image">
	<p class="product-name">Product Name</p>
	<p style="color: #292627; font-weight: bolder ; font-size:15px"> HURRAY! The product you are looking for is  PRICE
	<a style="background-color: #000; color: #fff; padding: 10px 20px; text-decoration: none; font-weight: bold; display: inline-block; margin-top: 10px; border-radius: 4px;" href="PRODUCT_URL">Place Order</a>
  </div>
</body>
</html>`

var emailClient *mail.Dialer

var systemEmailID = "rohith.knaidu0125@gmail.com" // default value

func initEmailClient() error {
	projectNumber := os.Getenv(envProjectNumber)

	if os.Getenv(envGmailID) != "" {
		systemEmailID = os.Getenv(envGmailID)
	}

	password := os.Getenv(envGmailPassword)
	if password == "" { // if password not in env var, get it from secrets manager
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
			Name: fmt.Sprintf("projects/%s/secrets/%s/versions/1", projectNumber, envGmailPassword),
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

func sendTrackNotificationEmail(t trackInput, p scrape.Product) error {
	m, err := prepareTrackNotificationEmail(t, p)
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

func prepareTrackNotificationEmail(t trackInput, p scrape.Product) (*mail.Message, error) {
	var emailBody string
	log.Println("creating mail")
	m := mail.NewMessage()
	m.SetHeader("From", systemEmailID)
	m.SetHeader("To", t.EmailID)
	if t.TypeOfRequest == requestTypeAvailability {
		m.SetHeader("Subject", "Availability update Notification")
	} else if t.TypeOfRequest == requestTypePrice {
		m.SetHeader("Subject", "price update Notification")
	} else {
		return nil, fmt.Errorf("invalid request type %s", t.TypeOfRequest)
	}
	m.Embed("./chrome-exten/logo.png")
	emailBody = strings.Replace(notificationEmailBody, "PRODUCT_URL", t.URL, -1)
	emailBody = strings.Replace(emailBody, "Product Name", p.Name, -1)
	emailBody = strings.Replace(emailBody, "product-image", p.Image, -1)
	if p.Price != 0 {
		priceString := strconv.FormatFloat(p.Price, 'f', -1, 64)
		emailBody = strings.Replace(emailBody, "PRICE", "available at ₹"+priceString+"</p>", -1)
	} else {
		emailBody = strings.Replace(emailBody, "PRICE", "available now</p>", -1)
	}

	m.SetBody("text/html", emailBody)
	return m, nil
}
