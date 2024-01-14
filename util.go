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

func validateAndCleanup(t *trackInput) error {

	if t.URL == "" {
		return fmt.Errorf("product url is mandatory field")
	}

	if t.TypeOfRequest == requestTypeAvailability ||
		t.TypeOfRequest == requestTypePrice {
		if t.EmailID == "" {
			return fmt.Errorf("email id is mandatory field")
		}
	}

	u, err := url.Parse(t.URL)
	if err != nil {
		return fmt.Errorf("error parsing url %s: %v", t.URL, err)
	}

	if _, ok := supportedWebsites[u.Hostname()]; !ok {
		return fmt.Errorf("url %s not supported: %w", t.URL, websiteNotSupported)
	}

	t.URL = scrape.CleanupURL(u).String()

	return nil
}
