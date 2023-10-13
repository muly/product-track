package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	scrape "github.com/muly/product-scrape"
)

// function for processing a url according the url provided
func callScraping(rawURL string) (scrape.Product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return scrape.Product{}, err
	}
	switch u.Hostname() {
	case "scrapeme.live":
		return scrape.GetScraper(u.Hostname())(rawURL)
	case "www.flipkart.com":
		return scrape.GetScraper(u.Hostname())(rawURL)
	case "www.amazon.in":
		return scrape.GetScraper(u.Hostname())(rawURL)
	case "localhost", "smuly-test-ground.ue.r.appspot.com":
		log.Println("scraping localhost")
		return integrationTestingMock(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return scrape.Product{}, fmt.Errorf("%s is not supported", u.Hostname())
	}
}

func integrationTestingMock(rawURL string) (scrape.Product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return scrape.Product{}, err
	}

	switch u.Path {
	case "/mock/amazon_available.html":
		return scrape.GetScraper("www.amazon.in")(rawURL)
	case "/mock/amazon_unavailable.html":
		return scrape.GetScraper("www.amazon.in")(rawURL)
	case "/mock/flipkart_available.html":
		return scrape.GetScraper("www.flipkart.com")(rawURL)
	case "/mock/flipkart_unavailable.html":
		return scrape.GetScraper("www.flipkart.com")(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Path)
		return scrape.Product{}, errors.New("unsupported URL path")
	}
}
