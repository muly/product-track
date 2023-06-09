package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func scrapeme(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		p.Price, err = checkPrice(h.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}

func flipkart(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("div._2JC05C", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	c.OnHTML("div._30jeq3._16Jk6d", func(h *colly.HTMLElement) {
		p.Price, err = checkPrice(h.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}

func amazon(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(h *colly.HTMLElement) {
		p.Price, err = checkPrice(h.ChildText("span.a-price-whole"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)
	p.Url = url

	return p, err
}
