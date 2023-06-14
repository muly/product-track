package main


import (
	"context"
	"fmt"
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
	ny := client.Collection("track_request")
	wr, err := ny.Parent.Create(ctx, trackInput{
		Url:           "www.youtube.com",
		TypeOfRequest: "priceRequest",
		MinThreshold:  1500,
	})
	if err != nil {
		log.Printf("error in create",err)
		return 
	}
	fmt.Println(wr)

}
