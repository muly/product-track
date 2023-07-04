package main

import (
	"context"
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
	if err := l.get(r.Context(), []filter{filter{"ProcessedDate", "<", todayDate}}); err != nil {
		log.Println("trackInputList.get() error:", err)
		return
	}
	log.Println("all records :::::::::::::", len(l), l)

	// make into batches
	var batch []trackInputList
	batch = append(batch, l) // TODO: need to split data into batches. for now only 1 batch

	ctx := r.Context()
	// go routine: process the batch.
	for _, b := range batch {
		processRequestBatch(ctx, b) // TODO: go routine resulting in context cancelled error
	}
}

func processRequestBatch(ctx context.Context, l trackInputList) {
	for _, t := range l {
		p, err := process(t.Url)
		if err != nil {
			log.Printf("error processing %s request for %s", t.TypeOfRequest, t.Url)
			updates := map[string]interface{}{
				"ProcessStatus": "ERROR",
				"ProcessNotes":  err.Error(),
			}
			if updateErr := t.patch(ctx, updates); updateErr != nil {
				log.Printf("Failed to update status field for document %s: %v\n", t.id(), updateErr)
			}

			continue
		}
		if shouldNotify(t, p) {
			if err := notify(t); err != nil {
				log.Printf("error sending notification: %s request for %s", t.TypeOfRequest, t.Url)
				updates := map[string]interface{}{
					"ProcessStatus": "ERROR",
					"ProcessNotes":  err.Error(),
				}
				if updateErr := t.patch(ctx, updates); updateErr != nil {
					log.Printf("Failed to update status field for document %s: %v\n", t.id(), updateErr)
				}

				continue
			}
		}

		updates := map[string]interface{}{
			"ProcessedDate": time.Now(),
			"ProcessStatus": "SUCCESS",
		}

		if err := t.patch(ctx, updates); err != nil {
			log.Printf("Failed to update processed_date and status fields for document %s: %v\n", t.id(), err)
		}

		// { // just for debugging
		// 	tOut := trackInput{Url: t.Url, TypeOfRequest: t.TypeOfRequest}
		// 	err = tOut.getByID(ctx)
		// 	if err != nil {
		// 		log.Printf("Failed  getByID(%v): %v\n", t.id(), err)
		// 		continue
		// 	}

		// 	log.Println("getByID output :::::::::::", tOut)
		// }
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
