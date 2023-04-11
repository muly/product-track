package main

import (
	"fmt"
	"net/url"


	"github.com/gocolly/colly"
)

type product struct {
	price, availabitlity string
}

func main() {

	rawURL := "https://www.amazon.com/Sphero-RVR-Programmable-Programmers-Customizable/dp/B0BLF8CLQF"
	


	 
	
	u,err := url.Parse(rawURL)
    

	if err != nil{
		// TODO:
	}

	p := product{}

	switch u.Hostname() {
	case "scrapeme.live":
		p, err = scrapeme(rawURL)
	case "ediblelandscaping.com":
		p,err=ediblelandscaping(rawURL)
	case "www.amazon.com":
        p, err=amazon(rawURL)
	default:
		fmt.Printf("%s is not supported\n", u.Hostname())
		return

	}

	if err != nil {
		// TODO:
	}

	fmt.Printf("%+v", p)
}

func scrapeme(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.availabitlity = h.Text
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
	var r product
	var err error

	c := colly.NewCollector()
	c.OnHTML("p.note", func(h *colly.HTMLElement) {
		r.availabitlity = h.Text
	})
	c.OnHTML("table.prices", func(h *colly.HTMLElement) {
	    r.price=h.ChildText("td")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s\n", r.URL))
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)

	return r, err
}

func amazon(url string) (product, error) {
	var s product
	var err error

	c := colly.NewCollector()
	c.OnHTML("", func(h *colly.HTMLElement) {
		s.availabitlity = h.Text
	})
	c.OnHTML("span.a-offscreen", func(h *colly.HTMLElement) {
	    s.price=h.ChildText(".a-price-whole")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s\n", r.URL))
	})
	c.OnError(func(r *colly.Response, scrapeErr error) {
		err = scrapeErr
	})
	c.Visit(url)

	return s, err
}

