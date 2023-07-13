package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// declaring variable at package level to access it easily
var firestoreClient *firestore.Client

// function for initializing firestore
func initFirestore(ctx context.Context) {
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		projectID = firestore.DetectProjectID
	}
	opt := option.WithCredentialsFile("path/to/credentials.json")

	var err error
	firestoreClient, err = firestore.NewClient(ctx, projectID,opt)
	if err != nil {
		log.Println("error occurred during database", err)
		return
	}
}
