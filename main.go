package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.portland5.com"),
	)

	c.OnHTML(".views-row", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".views-field-title .field-content a"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited", r.Request.URL)
	// })

	c.Visit("https://www.portland5.com/event-types/kids")
}
