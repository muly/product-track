package main

import (
	"fmt"

	"github.com/gocolly/colly"
)
type product struct{
	url,price,availabitlity string
}
func main(){
    
	c:=colly.NewCollector()
	c.OnHTML("li.product",func(h *colly.HTMLElement) {
		product:=product{}
		product.url=h.ChildAttr("a","href")
		product.price=h.ChildText(".price")
		 fmt.Println(product)
	
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s",r.URL))
	})
	c.OnError(func(r *colly.Response, err error) {
      fmt.Printf("error occured  %s\n",err.Error())
	})
	c.Visit("https://scrapeme.live/shop/")

}