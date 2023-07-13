package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type product struct {
	Url          string  `json:"url"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}

const requestTypeAvailability = "AVAILABILITY"
const requestTypePrice = "PRICE"

// function for checking availability using regular expressions
func checkAvailability(s string) bool {
	if strings.Contains(s, "In stock") {
		return true
	}
	hurryRegexList := []string{
		`^[Hh]urry, Only ([0-9]+) left!$`,
		`^Delivery by([0-9]+) [a-zA-Z], [a-zA-Z]$`, //Delivery by19 Jul, Wednesday|Free₹40?
	}
	for _, hurryRegex := range hurryRegexList {
		hurryRegex := regexp.MustCompile(hurryRegex)
		if hurryRegex.MatchString(s) {
			return true
		}
	}
	return false
}

// function for converting price from converting string type to float64  //price conversion
func priceConvertor(price string) (float64, error) {
	currencyList := []string{"₹", "$", "£"}
	price = strings.Replace(price, ",", "", -1)
	for _, c := range currencyList {
		price = strings.Replace(price, c, "", -1)
	}
	s, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("error occurred while parsing price", err)
		return 0, err
	}
	return s, nil
}

// function for conditions to satisfy for sending email   //notify condition
func notifyConditions(i trackInput, p product) bool {
	if i.TypeOfRequest == requestTypePrice && p.Price < i.MinThreshold {
		return true
	}
	if i.TypeOfRequest == requestTypeAvailability && p.Availability {
		return true
	}
	return false
}

// function for calling  sendemail function
func shouldNotify(t trackInput) error { //should notify
	sendEmail(t)
	return nil
}
