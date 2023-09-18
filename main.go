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
	router.GET("/execute-request", executeRequestHandler)
	router.POST("/store-email", storeEmailHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
	go staticServer();
}

func staticServer(){
	fs:=http.FileServer(http.Dir("./integration_testing/mock_websites/amazon_available.html"))
	http.Handle("/",fs)
	log.Printf("Listening on port:9090")
	err:=http.ListenAndServe(":9090",nil)
	if err !=nil {
		log.Fatal(err)
	}
}