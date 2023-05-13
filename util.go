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
	hurryRegex := regexp.MustCompile(`^[Hh]urry only ([0-9]+) items left!$`)

	if strings.Contains(s, "In stock") {
		return true
	} else if hurryRegex.MatchString(s) {
		return true
	}

	return false
}

func checkPrice(price string) (float64, error) {
	price = strings.Replace(price, ",", "", -1)
	s, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("error occurred while parsing value", err)
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
