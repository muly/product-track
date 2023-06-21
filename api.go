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
	Url           string  `json:"url"`
	MinThreshold  float64 `json:"min_threshold"`
	TypeOfRequest string  `json:"type_of_request"`
	activeRequest bool
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
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during handling the url", err)
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
	ctx := r.Context()
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during handling the url", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.TypeOfRequest = requestTypeAvailability

<<<<<<< HEAD
	id := fmt.Sprintf("[%s][%s]", url.QueryEscape(t.Url),requestTypeAvailability)

	_, err := client.Collection("track_requests").Doc(id).Set(ctx, t)
=======
	_, err := client.Collection("track_requests").Doc(t.id()).Set(ctx, t)
>>>>>>> c27568400b9b0f6c5789cce81c1c47b1f823cafb
	if err != nil {
		log.Println("error during firestore write", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore write ", err)))
		return
	}

	var tOut trackInput
	if err := tOut.getByID(ctx); err != nil {
		log.Println("error during firestore get", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore get ", err)))
		return
	}

	log.Printf("data retrieved from firestore %+v\n", tOut)
}

func priceHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during price  handling ", err)
		// TODO: return with status code 400
	}
	fmt.Println(t)
	//TODO:need to persist the request in a database
}
