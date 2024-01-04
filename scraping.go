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
	var scrapeFunctions []func(url string) (scrape.Product, error)
	if _, ok := supportedWebsites[u.Hostname()]; ok {
		scrapeFunctions = scrape.GetProductFunctions(u.Hostname())
	} else if u.Hostname() == "localhost" || u.Hostname() == "smuly-test-ground.ue.r.appspot.com" {
		log.Println("scraping for integration testing")
		return integrationTestingMock(rawURL)
	} else {
		log.Printf("%s is not supported\n", u.Hostname())
		return scrape.Product{}, fmt.Errorf("%s is not supported", u.Hostname())
	}

	var p scrape.Product
	for _, f := range scrapeFunctions {
		p, err = f(rawURL)
		if err != nil {
			// TODO: log
			continue
		}
		break
	}
	return p, err
}

func integrationTestingMock(rawURL string) (scrape.Product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return scrape.Product{}, err
	}

	switch u.Path {
	case "/mock/amazon_available.html":
		return scrape.GetProductFunctions("www.amazon.in")[0](rawURL)
	case "/mock/amazon_unavailable.html":
		return scrape.GetProductFunctions("www.amazon.in")[0](rawURL)
	case "/mock/flipkart_available.html":
		return scrape.GetProductFunctions("www.flipkart.com")[0](rawURL)
	case "/mock/flipkart_unavailable.html":
		return scrape.GetProductFunctions("www.flipkart.com")[0](rawURL)
	default:
		log.Printf("%s is not supported\n", u.Path)
		return scrape.Product{}, errors.New("unsupported URL path")
	}
}
