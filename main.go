package main

import (
	"fmt"

	"net/url"
)

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
