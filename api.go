package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type trackInput struct {
	Url string `json:"url"`
}
type priceTrackInput struct {
	Url          string  `json:"url"`
	MinThreshold float64 `json:"minThreshold"`
}

func handleRequest() {
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/products", productss)
	// log.Fatal(http.ListenAndServe(":8006", nil))
	router := httprouter.New()
	router.POST("/track/availability", availabilityHandler)
	router.POST("/product", productHandler)
	router.POST("/track/price", priceHandler)
	log.Fatal(http.ListenAndServe(":8006", router))
}

func productHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal("error during handling the url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pr, err := process(t.Url)
	if err != nil {
		log.Fatal("error in processing url", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(pr)
	if err != nil {
		log.Fatal("error in encoding product", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func availabilityHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal("error during handling the url", err)
	}
	fmt.Println(t.Url)

}
func priceHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var t priceTrackInput
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal("error during price  handling ", err)
	}
	fmt.Println(t)
}
