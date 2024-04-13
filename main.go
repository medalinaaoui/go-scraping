package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Movie struct {
	Name string `json:"name"`
	Url string `json:"url"`
}


func main() {
	c := colly.NewCollector()
	c.OnHTML(".entry-box .entry-image a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Visiting:", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("[data-quality='5']", func(e *colly.HTMLElement) {
		// Scrape data here
	})


	c.Visit("https://ak.sv/search?q=the+irishman")
		fmt.Println("hey")

}