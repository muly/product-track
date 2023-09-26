package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type trackInput struct {
	Url           string  `json:"url"`
	MinThreshold  float64 `json:"min_threshold"`
	TypeOfRequest string  `json:"type_of_request"`
	ProcessedDate time.Time
	ProcessStatus string
	EmailId       string `json:"emailid"`
}

const (
	fieldProcessStatus = "ProcessStatus"
	fieldProcessedDate = "ProcessedDate"
	fieldProcessNotes  = "ProcessNotes"

	processStatusSuccess = "SUCCESS"
	processStatusError   = "ERROR"
)

type User struct {
	Email string `json:"email"`
}

func storeEmailHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u User
	log.Println("store email end point is started")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := u.upsert(r.Context()); err != nil {
		log.Println("error during firestore ups func", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore ups func", err)))
		return
	}
}

// api function for  execute_request  end point
func executeRequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	todayDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	// get records
	var lPrev trackInputList
	filters := []filter{
		{"ProcessedDate", "<", todayDate},
	}
	if err := lPrev.get(r.Context(), filters); err != nil {
		log.Println("trackInputList.get() [old records] error:", err)
		return
	}

	var lTodayFailed trackInputList
	filters = []filter{
		{"ProcessedDate", ">", todayDate.Add(-time.Second)},
		{"ProcessStatus", "==", processStatusError},
	}
	if err := lTodayFailed.get(r.Context(), filters); err != nil {
		log.Println("trackInputList.get() [today failed records] error:", err)
		return
	}

	var l trackInputList = append(lPrev, lTodayFailed...)
	log.Println("total records to be processed", len(l))

	// make into batches
	var batch []trackInputList
	batch = append(batch, l) // TODO: need to split data into batches. for now only 1 batch

	ctx := r.Context()

	var pl patchList
	// go routine: process the batch.
	for _, b := range batch {
		l := processRequestBatch(ctx, b) // TODO: go routine resulting in context cancelled error
		pl = append(pl, l...)
	}

	pl.patch(ctx)
}

// processRequestBatch processes the given batch of track inputs and return the payload for process updates back to database
func processRequestBatch(ctx context.Context, l trackInputList) patchList {
	processNotes := ""
	var updatesTodo patchList
	for _, t := range l {
		p, err := callScraping(t.Url)
		if err != nil {
			log.Printf("error processing %s request for %s: %v", t.TypeOfRequest, t.Url, err)
			processNotes = "scrape error: " + err.Error()
			updatesTodo = append(updatesTodo, patch{
				typeOfRequest: t.TypeOfRequest,
				emailid:       t.EmailId,
				url:           t.Url,
				patchData: map[string]interface{}{
					fieldProcessedDate: time.Now(),
					fieldProcessStatus: processStatusError,
					fieldProcessNotes:  processNotes,
				}})
			continue
		}

		if shouldNotify(t, p) {
			if err := sendTrackNotificationEmail(t); err != nil {
				log.Printf("error sending notification: %s request for %s", t.TypeOfRequest, t.Url)
				updatesTodo = append(updatesTodo, patch{
					typeOfRequest: t.TypeOfRequest,
					url:           t.Url,
					patchData: map[string]interface{}{
						fieldProcessedDate: time.Now(),
						fieldProcessStatus: processStatusError,
						fieldProcessNotes:  err.Error(),
					}})
				continue
			}

			updatesTodo = append(updatesTodo, patch{
				typeOfRequest: t.TypeOfRequest,
				emailid:       t.EmailId,
				url:           t.Url,
				patchData: map[string]interface{}{
					fieldProcessedDate: time.Now(),
					fieldProcessStatus: processStatusSuccess,
					fieldProcessNotes:  "processed & notification sent",
				}})
		} else {
			updatesTodo = append(updatesTodo, patch{
				typeOfRequest: t.TypeOfRequest,
				emailid:       t.EmailId,
				url:           t.Url,
				patchData: map[string]interface{}{
					fieldProcessedDate: time.Now(),
					fieldProcessStatus: processStatusSuccess,
					fieldProcessNotes:  "processed",
				}})
		}

	}
	return updatesTodo
}

// function for processing the url
func productHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t trackInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println("error during handling the url", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pr, err := callScraping(t.Url)
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

// function for availability request with /track/availability end point
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
}

// function for price request with /track/price end point
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
