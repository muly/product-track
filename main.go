package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {
	handleRequest()
	ctx:=context.Background()
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		projectID = firestore.DetectProjectID
	}
	var err error
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("error occurred during database", err)
	}

}

func process(rawURL string) (product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return product{}, err
	}
	switch u.Hostname() {
	case "scrapeme.live":
		return scrapeme(rawURL)
	case "www.flipkart.com":
		return flipkart(rawURL)
	case "www.amazon.in":
		return amazon(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
	}
}
