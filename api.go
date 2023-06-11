package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type trackInput struct {
	Url string `json:"url"`
}
type priceTrackInput struct {
	Url          string  `json:"url"`
	MinThreshold float64 `json:"min_threshold"`
}

func handleRequest() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8006"
		log.Printf("Defaulting to port %s", port)
	}
	router := httprouter.New()
	router.POST("/track/availability", availabilityHandler)
	router.POST("/product", productHandler)
	router.POST("/track/price", priceHandler)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func productHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t);err !=nil {
		log.Fatal("error during handling the url",err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pr, err := process(t.Url)
	if err != nil {
		log.Println("error in processing url", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(pr)
	if err != nil {
		log.Println("error in encoding product", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func availabilityHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t);err != nil {
		log.Println("error during handling the url", err)
	}
	fmt.Println(t.Url) 
	//TODO:need to persist the request in a database
}

func priceHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t priceTrackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t);err != nil {
		log.Println("error during price  handling ", err)
	}
	fmt.Println(t)
	//TODO:need to persist the request in a database
}
