package main

import (
	"fmt"

	"github.com/gocolly/colly"
)
type product struct{
	price,availabitlity string
}
func main(){
    var p product
	c:=colly.NewCollector()
	c.OnHTML("p.stock.in-stock",func(h *colly.HTMLElement) {
		p.availabitlity=h.Text
	

	})
	c.OnHTML("p.price",func(h *colly.HTMLElement) {
	  p.price=h.Text
	  
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("visiting %s",r.URL))
	})
	c.OnError(func(r *colly.Response, err error) {
      fmt.Printf("error occured  %s\n",err.Error())
	})
	c.Visit("https://scrapeme.live/shop/Charmander/")
	fmt.Println(p)

}