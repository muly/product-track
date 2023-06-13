package main

import (
	"log"
	"net/url"
)

func main() {
	handleRequest()
}

func process(rawURL string) (product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return product{}, err
	}
	switch u.Hostname() {
	case "scrapeme.live":
		return scrapeme(rawURL)
	case "www.flipkart.com":
		return flipkart(rawURL)
	case "www.amazon.in":
		return amazon(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
	}
}
