package main

import (
	"fmt"
	"net/url"


	"github.com/gocolly/colly"
)

type product struct {
	price, availability string
}

func main() {

	rawURL := "https://www.amazon.in/Noise-Launched-Bluetooth-Calling-Tracking/dp/B0BJ72WZQ7/ref=sr_1_2?pd_rd_r=b244f48d-24db-4f55-8fe7-20bbe6ade0d2&pd_rd_w=Iq0zo&pd_rd_wg=EZ5WS&pf_rd_p=3a59b28c-4626-48f9-b66f-114571ee563d&pf_rd_r=XA5YK79N3KZGRAMATR6X&qid=1681233130&sr=8-2&th=1"
	u,err := url.Parse(rawURL)
    if err != nil{
		fmt.Println("error occurred ",err)
	}
	p := product{}
    switch u.Hostname() {
	case "scrapeme.live":
		p , err = scrapeme(rawURL)
	case "ediblelandscaping.com":
		p ,err=ediblelandscaping(rawURL)
	case "www.amazon.in":
        p , err=amazon(rawURL)
	default:
		fmt.Printf("%s is not supported\n", u.Hostname())
		return
    }
    if err != nil {
		fmt.Println("error occurred",err)
	}

	fmt.Printf("%+v", p)
}

func scrapeme(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.availability = h.Text
	})
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		p.price = h.Text
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s\n", r.URL))
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)

	return p, err
}
func ediblelandscaping(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("p.note", func(h *colly.HTMLElement) {
		p.availability = h.Text
	})
	c.OnHTML("table.prices", func(h *colly.HTMLElement) {
	    p.price=h.ChildText("td")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s\n", r.URL))
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)

	return p, err
}

func amazon(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
		p.availability = h.Text
	})
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center",func(h *colly.HTMLElement) {
		p.price=h.ChildText("span.a-price-whole")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s\n", r.URL))
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)

	return p, err
}

