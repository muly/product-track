package main

import (
	"fmt"
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
func checkAvailability( s string)(bool ){
	if strings.Contains(s,"In stock"){
		return true
	} else if s == "Hurry, Only %d left!" {
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				return  true
			}
		}
	} 
	return false

}
func checkPrice(price string)(float64){
	var f float64
	var j float64
	j = 1
	var i int
	for i = 0; i < len(price); i++ {
		if price[i] >= '0' && price[i] <= '9' {
			//fmt.Printf("%T",price[i]-48)
			f = (f * 10) + float64(price[i]-48)
		}
		if price[i] == '.' {
			break
		}
	}
	for i < len(price) {
		if price[i] >= '0' && price[i] <= '9' {
			j = j * 10
			f = f + float64(price[i]-48)/j

		}
		i++
	}

	return f

}

func scrapeme(url string) (product, error) {
	var p product
	var err error
	c := colly.NewCollector()
	c.OnHTML("p.stock.in-stock", func(h *colly.HTMLElement) {
		p.availability=checkAvailability(h.Text)
	
     })
	c.OnHTML("p.price", func(h *colly.HTMLElement) {
		price := h.Text
		var f float64
		var j float64
		j = 1
		var i int
		for i = 0; i < len(price); i++ {
			if price[i] >= '0' && price[i] <= '9' {
				//fmt.Printf("%T",price[i]-48)
				f = (f * 10) + float64(price[i]-48)
			}
			if price[i] == '.' {
				break
			}
		}
		for i < len(price) {
			if price[i] >= '0' && price[i] <= '9' {
				j = j * 10
				f = f + float64(price[i]-48)/j

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
func flipkart(url string) (product, error) {
	var p product
	var err error
	var statement string
	c := colly.NewCollector()
	c.OnHTML("div._2JC05C", func(h *colly.HTMLElement) {

		if statement == "                                                In stock                               " {
			p.availability = true
		} else if statement == "Hurry, Only %d left!" {
			for i := 0; i < len(statement); i++ {
				if statement[i] >= '0' && statement[i] <= '9' {
					p.availability = true
				}
			}
		} else {
			p.availability = false
		}

	})
	c.OnHTML("div._30jeq3._16Jk6d", func(h *colly.HTMLElement) {
		price := h.Text
		var f float64
		var j float64
		j = 1
		var i int
		for i = 0; i < len(price); i++ {
			if price[i] >= '0' && price[i] <= '9' {
				//fmt.Printf("%T",price[i]-48)
				f = (f * 10) + float64(price[i]-48)
			}
			if price[i] == '.' {
				break
			}
		}
		for i < len(price) {
			if price[i] >= '0' && price[i] <= '9' {
				j = j * 10
				f = f + float64(price[i]-48)/j

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

func amazon(url string) (product, error) {
	var p product
	var err error

	c := colly.NewCollector()
	c.OnHTML("#availability", func(h *colly.HTMLElement) {
       p.availability=checkAvailability(h.Text)
	})
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(h *colly.HTMLElement) {
		p.price =checkPrice(h.ChildText("span.a-price-whole"))
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

func shouldNotify(i input)(bool,error){
	var err error
	var p product 
	var r input
	flag:=1
	if p.price <= r.minThreshold{
       flag=0 
	}
	if flag==0{
		return true,err
	}else {
		return false,err
	}
}
