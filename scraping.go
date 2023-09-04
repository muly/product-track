package main

import (
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
	case "www.amazon.in","localhost:5500":
		return amazon(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
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
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
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
