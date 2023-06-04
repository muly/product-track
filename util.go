package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type product struct {
	price        float64   
	availability bool
}
type input struct {
	typeOfRequest string
	minThreshold  float64
}

const requestTypeAvailability = "AVAILABILITY"
const requestTypePrice = "PRICE"

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

func checkPrice(price string) (float64, error) {
	currencyList := []string{"₹", "$"}
	price = strings.Replace(price, ",", "", -1)
	for _, c:= range currencyList{
		price = strings.Replace(price, c, "", -1)
	}
	s, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("error occurred while parsing price", err)
		return 0, err
	}
	return s, nil
}

func shouldNotify(i input, p product) bool {
	if i.typeOfRequest == requestTypePrice && p.price < i.minThreshold {
		return true
	}
	if i.typeOfRequest == requestTypeAvailability && p.availability {
		return true
	}
	return false
}
