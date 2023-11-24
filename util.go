package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	scrape "github.com/muly/product-scrape"
)

const (
	requestTypeAvailability = "AVAILABILITY"
	requestTypePrice        = "PRICE"
)

var websiteNotSupported error = fmt.Errorf("unsupported website")

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

func readSupportedWebsites() (map[string]bool, error) {
	f, err := os.Open("supported_websites")
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	output := make(map[string]bool)

	for s.Scan() {
		output[s.Text()] = true
	}

	return output, nil
}

func validate(t trackInput) error {
	if t.URL == "" || t.EmailID == "" {
		return fmt.Errorf("some of the mandatory fields are missing")
	}

	u, err := url.Parse(t.URL)
	if err != nil {
		return fmt.Errorf("error parsing url %s: %v: %w", t.URL, err, websiteNotSupported)
	}

	if _, ok := supportedWebsites[u.Hostname()]; !ok {
		return fmt.Errorf("url %s not supported: %w", t.URL, websiteNotSupported)
	}

	return nil
}
