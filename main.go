package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {

	if err := initFirestore(context.Background()); err != nil {
		log.Printf("failed to create firestore client: %v", err)
		os.Exit(1)
	}

	if err := initEmailClient(); err != nil {
		log.Printf("failed to create email client: %v", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8006"
		log.Printf("Defaulting to port %s", port)
	}
	router := httprouter.New()
	router.POST("/track/availability", availabilityHandler)
	router.POST("/product", productHandler)
	router.POST("/track/price", priceHandler)
	router.POST("/execute-request", executeRequest)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
