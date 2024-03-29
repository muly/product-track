package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var supportedWebsites = make(map[string]bool)

func main() {
	fmt.Printf("Git Commit Hash: %s\n", os.Getenv("COMMIT_HASH"))

	if err := initFirestore(context.Background()); err != nil {
		log.Printf("failed to create firestore client: %v", err)
		os.Exit(2)
	}

	if err := initEmailClient(); err != nil {
		log.Printf("failed to create email client: %v", err)
		os.Exit(3)
	}

	var err error
	supportedWebsites, err = readSupportedWebsites()
	if err != nil {
		log.Printf("failed to read supported websites list: %v", err)
		os.Exit(4)
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
	router.GET("/mock/*path", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		path := ps.ByName("path")
		http.ServeFile(w, r, "./integration_testing/mock_websites/"+path)
	})

	router.GET("/privacy-policy.html", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "./chrome-exten/privacy-policy.html")
	})
	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, router))

}
