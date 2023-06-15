package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/julienschmidt/httprouter"
)

// type trackInput struct {
// 	Url string `json:"url"`
// }

type trackInput struct {
	Url           string  `json:"url"`
	MinThreshold  float64 `json:"min_threshold"`
	TypeOfRequest string  `json:"type_of_request"`
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
		// TODO: return with status code 400
		return
	}
	fmt.Println(t.Url)
	//TODO:need to persist the request in a database

	id := fmt.Sprintf("%s|%s", url.QueryEscape(t.Url), t.TypeOfRequest)
	// id = "test123"
	fmt.Println("ID ###############", id)
	fmt.Println("client ###############", client)

	dref := client.Doc(id)
	fmt.Println("dref ###############", dref)

	_, err := dref.Set(ctx, t)
	if err != nil {
		log.Println("error during firestore write", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore write", err)))
		return
	}

	d, err := client.Doc(id).Get(ctx)
	if err != nil {
		log.Println("error during firestore get", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore get", err)))
		return
	}
	var out trackInput
	if err := d.DataTo(&out); err != nil {
		log.Println("error during firestore datato", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during firestore datato", err)))
		return
	}

	fmt.Println("data retrieved from firestore", out)

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
