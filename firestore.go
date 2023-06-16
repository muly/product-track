package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

func clientSide(ctx context.Context) {
	log.Println("clientside is started")
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		projectID = firestore.DetectProjectID
	}
	var err error
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Println("error occurred during database", err)
		return
	}

	// ny := client.Collection("track_request")
	// log.Println("writing document is started")
	// wr, err := ny.Parent.Create(ctx, trackInput{
	// 	Url:           "www.youtube.com",
	// 	TypeOfRequest: "priceRequest",
	// 	MinThreshold:  1500,
	// })
	// if err != nil {
	// 	log.Printf("error in create", err)
	// 	return
	// }
	// fmt.Println(wr)

}
