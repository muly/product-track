package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type trackInput struct {
	Url           string  `json:"url"`
	MinThreshold  float64 `json:"min_threshold"`
	TypeOfRequest string  `json:"type_of_request"`
	ProcessedDate time.Time
	ProcessStatus string
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
	router.POST("/execute-request", executeRequest)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func executeRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get records
	todayDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	var l trackInputList
	if err := l.get(r.Context(), []filter{filter{"ProcessedDate", ">", todayDate}}); err != nil {
		log.Println("trackInputList.get() error:", err)
		return
	}
	log.Println("all records :::::::::::::", len(l), l)

	// make into batches
	var batch []trackInputList
	batch = append(batch, l) // TODO: need to split data into batches. for now only 1 batch

	for _, b := range batch {
		go processRequestBatch(b)
	}

	// go routine: process the batch.

}

func processRequestBatch(l trackInputList) {
	for _, t := range l {
		p, err := process(t.Url)
		if err != nil {
			log.Println("error processing %s request for %s", t.TypeOfRequest, t.Url)
			// TODO: also update the track_request table in a status field.
			continue
		}
		if shouldNotify(t, p) {
			if err := notify(t); err != nil {
				log.Println("error sending notification: %s request for %s", t.TypeOfRequest, t.Url)
				// TODO: also update the track_request table in a status field.
				continue
			}
		}
		// TODO: update the records processed_date field with current timestamp, and status field as SUCCESS
	}
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
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during handling the url", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.TypeOfRequest = requestTypeAvailability
	if err := t.upsert(r.Context()); err != nil {
		log.Println("error during firestore upsert in availability handler", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore upsert in availability handler", err)))
		return
	}

	// // TODO: remove this after testing

}

func priceHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during price  handling ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.TypeOfRequest = requestTypePrice
	if err := t.upsert(r.Context()); err != nil {
		log.Println("error during firestore upsert in availability handler", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore upsert in availability handler", err)))
		return
	}
}
