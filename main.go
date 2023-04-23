package main

import (
	"fmt"
	"regexp"
	"strconv"

	"strings"

	"net/url"

	"github.com/gocolly/colly"
)

type product struct {
	price        float64
	availability bool
}
type input struct {
	url           string
	typeOfRequest string
	minThreshold  float64
}

const requestTypeAvailability = "AVAILABILITY"
const requestTypePrice = "PRICE"

func main() {

	rawURL := "https://www.amazon.in/Bassbuds-Duo-Headphones-Water-Resistant-Assistance/dp/B09DD9SX9Z/ref=sr_1_1?_encoding=UTF8&_ref=dlx_gate_sd_dcl_tlt_a1f4109b_dt&content-id=amzn1.sym.0d1fafce-0d80-4ffc-b8c3-74f55ca06beb&pd_rd_r=c7d70bbf-7c14-4bf2-b7bf-3b4e92859c5e&pd_rd_w=MLpI8&pd_rd_wg=TZMUW&pf_rd_p=0d1fafce-0d80-4ffc-b8c3-74f55ca06beb&pf_rd_r=DZ2A8Y82Z4JPJE5B7A4W&qid=1682187704&sr=8-1&th=1"
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("invalid url ", err)
		return
	}
	p := product{}

	switch u.Hostname() {
	case "scrapeme.live":
		p, err = scrapeme(rawURL)
	case "flipkart.com":
		p, err = flipkart(rawURL)
	case "www.amazon.in":
		p, err = amazon(rawURL)
	default:
		fmt.Printf("%s is not supported\n", u.Hostname())
		return
	}
	if err != nil {
		fmt.Println("error occurred while scraping", err)
		return
	}

	fmt.Printf("%+v", p)

}
func checkAvailability(s string) bool {
	var r = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	if strings.Contains(s, "In stock") {
		return true
	} else if r.MatchString(s) {
		return true
	}

	return false

}
func checkPrice(price string) (float64,error) {
	price = strings.Replace(price, ",", "", -1)
	s, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("error occurred while parsing value", err)
		return 0,err
	}
	return s,nil
}

func scrapeme(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.availability = checkAvailability(h.Text)

	})
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		p.price,err= checkPrice(h.Text)

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
func flipkart(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("div._2JC05C", func(h *colly.HTMLElement) {
		p.availability = checkAvailability(h.Text)

	})
	c.OnHTML("div._30jeq3._16Jk6d", func(h *colly.HTMLElement) {
		p.price,err = checkPrice(h.Text)

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
		p.availability = checkAvailability(h.Text)
	})
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(h *colly.HTMLElement) {
		p.price,err = checkPrice(h.ChildText("span.a-price-whole"))
		
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

/*
  check availability-p.availability true == yes
  check price -if
*/

func shouldNotify(i input, p product) bool {

	if i.typeOfRequest == requestTypePrice && p.price < i.minThreshold {
		return true
	}
	if i.typeOfRequest == requestTypeAvailability && p.availability {
		return true
	}
	return false
}
