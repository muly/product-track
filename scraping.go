package main

import (
	"errors"
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

// function for processing a url according the url provided
func callScraping(rawURL string) (product, error) {
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
	case "localhost", "smuly-test-ground.ue.r.appspot.com":
		log.Println("scraping localhost")
		return integrationTestingMock(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
	}
}

func integrationTestingMock(rawURL string) (product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return product{}, err
	}

	path := u.Path

	switch path {
	case "/mock/amazon_available.html":
		return amazon(rawURL)
	case "/mock/amazon_unavailable.html":
		return amazon(rawURL)
	default:
		log.Printf("%s is not supported\n", path)
		return product{}, errors.New("unsupported URL path")
	}
}

// scraping function for collecting  scrapeme data
func scrapeme(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		p.Price, err = priceConvertor(h.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}

// scraping function for collecting  flipkart data
func flipkart(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("div._3XINqE", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	c.OnHTML("div._30jeq3._16Jk6d", func(h *colly.HTMLElement) {
		p.Price, err = priceConvertor(h.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}

// scraping function for collecting  amazon data
func amazon(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	//availability
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	//a-section.a-spacing-none.aok-align-center
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(h *colly.HTMLElement) {
		p.Price, err = priceConvertor(h.ChildText("span.a-price-whole"))
	})
	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}
