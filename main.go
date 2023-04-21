package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
)

type product struct {
	price        float64
	availability bool
}
type input struct {
	url           string
	typeOfRequest string
	minThreshold  float32
}

func main() {

	rawURL := "https://www.amazon.in/Noise-Launched-Bluetooth-Calling-Tracking/dp/B0BJ72WZQ7/ref=sr_1_2?pd_rd_r=b244f48d-24db-4f55-8fe7-20bbe6ade0d2&pd_rd_w=Iq0zo&pd_rd_wg=EZ5WS&pf_rd_p=3a59b28c-4626-48f9-b66f-114571ee563d&pf_rd_r=XA5YK79N3KZGRAMATR6X&qid=1681233130&sr=8-2&th=1"
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("error occurred ", err)
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
		fmt.Println("error occurred", err)
	}

	fmt.Printf("%+v", p)

}

func scrapeme(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		availability := h.Text
		if availability == "In stock" {
			p.availability = true
		} else {
			p.availability = false
		}

	})
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		price := h.Text
		s, err := strconv.ParseFloat(price,10)
		if err != nil {
			log.Printf("error")
		}
		p.price = s

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
		availability := h.Text
		if availability == " In stock " {
			p.availability = true
		} else {
			p.availability = false
		}
	})
	c.OnHTML("div._30jeq3._16Jk6d", func(h *colly.HTMLElement) {
		price := h.Text
		s, err := strconv.ParseFloat(price,10)
		if err != nil {
			log.Printf("error")
		}
		p.price = s
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
    var statement string
	c := colly.NewCollector()
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
		statement = h.Text
		fmt.Println(statement)
		if statement == "                                                In stock                               " {
			p.availability = true
		} else {
			p.availability = false
		}

	})
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(h *colly.HTMLElement) {
		price := h.ChildText("span.a-price-whole")
        var f float64
		var j float64
		j=1
		var i int
		for i=0 ;i<len(price);i++ {
			if price[i]>='0' && price[i]<='9'{
				//fmt.Printf("%T",price[i]-48)
                f=(f*10 ) + float64(price[i]-48)	
			}
			if price[i]=='.'{
               break
			}
		}
		for i<len(price){
			if price[i]>='0' && price[i]<='9'{
			  j=j*10
              f=f+float64(price[i]-48)/j

			}
		 i++	
		}


		p.price = f
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

//func shouldNotify(i input)(bool,error){
//var err error
//var p product
//if err!=nil{
//fmt.Println(err)
//}

//return true,err
//}
