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

//function for checking availability using regular expressions
func checkAvailability(s string) bool {
	if strings.Contains(s, "In stock") {
		return true
	}
	hurryRegexList := []string{
		`^[Hh]urry, only ([0-9]+) items left!$`,
		`^[Hh]urry only ([0-9]+) items left!$`,
	}
	for _, hurryRegex := range hurryRegexList {
		hurryRegex := regexp.MustCompile(hurryRegex)

		if hurryRegex.MatchString(s) {
			return true
		}
	}
	return false
}

//function for converting price from converting string type to float64 
func checkPrice(price string) (float64, error) {
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

//function for conditions to satisfy for sending email
func shouldNotify(i trackInput, p product) bool {
	if i.TypeOfRequest == requestTypePrice && p.Price < i.MinThreshold {
		return true
	}
	if i.TypeOfRequest == requestTypeAvailability && p.Availability {
		return true
	}
	return false
}

//function for calling  sendemail function
func notify(t trackInput) error {
	if err:=sendEmail();err!=nil{
		log.Println("error occured in notify function",err)
		return nil 
	}
	return nil
}
