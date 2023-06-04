package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Inputs []input

func pro(w http.ResponseWriter, r *http.Request) {
	inputs := Inputs{
		input{typeOfRequest: "price request", minThreshold: 1500},
	}
	fmt.Println("all products end point")
	json.NewEncoder(w).Encode(inputs)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " home page end point")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/input", pro)
	log.Fatal(http.ListenAndServe(":8006", nil))
}