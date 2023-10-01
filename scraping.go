package main

import (
	"errors"
	"fmt"
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
		return product{}, fmt.Errorf("%s is not supported", u.Hostname())
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
	case "/mock/flipkart_available.html":
		return flipkart(rawURL)
	case "/mock/flipkart_unavailable.html":
		return flipkart((rawURL))
	default:
		log.Printf("%s is not supported\n", path)
		return product{}, errors.New("unsupported URL path")
	}
}

type scrapeTags struct {
	availability string
	price        string
	priceChild   string
}

// generic scrape
func scrape(url string, t scrapeTags) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	//availability
	c.OnHTML(t.availability, func(h *colly.HTMLElement) {
		p.Availability = checkAvailability(h.Text)
	})
	//a-section.a-spacing-none.aok-align-center
	c.OnHTML(t.price, func(h *colly.HTMLElement) {
		price := h.Text
		if t.priceChild != "" {
			price = h.ChildText(t.priceChild)
		}
		p.Price, err = priceConvertor(price)
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

// scraping function for collecting  scrapeme data
func scrapeme(url string) (product, error) {
	scrapemeTags := scrapeTags{
		availability: "p.stock.in-stock",
		price:        "p.price",
		priceChild:   "",
	}
	return scrape(url, scrapemeTags)
}

// scraping function for collecting  flipkart data
func flipkart(url string) (product, error) {
	flipkartTags := scrapeTags{
		availability: "div._3XINqE",
		price:        "div._30jeq3._16Jk6d",
		priceChild:   "",
	}

	return scrape(url, flipkartTags)
}

// scraping function for collecting  amazon data
func amazon(url string) (product, error) {
	amazonTags := scrapeTags{
		availability: "#availability",
		price:        "div.a-section.a-spacing-none.aok-align-center",
		priceChild:   "span.a-price-whole",
	}
	return scrape(url, amazonTags)
}
