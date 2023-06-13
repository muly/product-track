package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

func clientSide(ctx context.Context) {
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
