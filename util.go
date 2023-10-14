package main

import scrape "github.com/muly/product-scrape"

const requestTypeAvailability = "AVAILABILITY"
const requestTypePrice = "PRICE"

// function for conditions to satisfy for sending email   //notify condition
func shouldNotify(i trackInput, p scrape.Product) bool {
	if (p == scrape.Product{}) {
		return false
	}
	if i.TypeOfRequest == requestTypePrice && p.Price < i.MinThreshold {
		return true
	}
	if i.TypeOfRequest == requestTypeAvailability && p.Availability {
		return true
	}
	return false
}
