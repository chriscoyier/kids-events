package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const (
	PORTLAND5_DOMAIN      = "www.portland5.com"
	PORTLAND5_LINK_PREFIX = "https://" + PORTLAND5_DOMAIN
	PORTLAND5_CRAWL_URL   = "https://www.portland5.com/event-types/kids"

	SELLWOOD_DOMAIN      = "www.sellwoodcommunityhouse.org"
	SELLWOOD_LINK_PREFIX = "https://" + SELLWOOD_DOMAIN
	SELLWOOD_CRAWL_URL   = "https://www.sellwoodcommunityhouse.org/community-events-1"

	ZOO_DOMAIN      = "www.oregonzoo.org"
	ZOO_LINK_PREFIX = "https://" + ZOO_DOMAIN
	ZOO_CRAWL_URL   = "https://www.oregonzoo.org/events"
)

func getPortland5Data() []KidsEvent {
	kidsEvents := []KidsEvent{}

	c := colly.NewCollector(
		colly.AllowedDomains(PORTLAND5_DOMAIN),
	)

	c.OnHTML(".views-row", func(e *colly.HTMLElement) {
		titleText := e.ChildText(".views-field-title .field-content a")
		url := PORTLAND5_LINK_PREFIX + e.ChildAttr(".views-field-title .field-content a", "href")
		date := e.ChildAttr(".date-display-single", "content")
		venue := e.ChildText(".views-field-field-event-venue")

		if titleText != "" {
			kidsEvents = append(kidsEvents, KidsEvent{
				Title:   titleText,
				URL:     url,
				Date:    date,
				Display: true,
				Venue:   venue,
			})
		}
	})

	c.Visit(PORTLAND5_CRAWL_URL)

	return kidsEvents
}

func getSellwoodData() []KidsEvent {
	kidsEvents := []KidsEvent{}

	c := colly.NewCollector(
		colly.AllowedDomains(SELLWOOD_DOMAIN),
	)

	c.OnHTML(".ProductList-item", func(e *colly.HTMLElement) {
		titleText := e.ChildText(".ProductList-title")
		url := SELLWOOD_LINK_PREFIX + e.ChildAttr(".ProductList-item-link", "href")

		// either need to get the data out of the title or follow the link to the next page and get it from there.
		date := "2099-01-01 00:00:00"

		venue := "Sellwood Community House"

		fmt.Println("SELLWOOD EVENT", titleText, url, date, venue)

		if titleText != "" {
			kidsEvents = append(kidsEvents, KidsEvent{
				Title:   titleText,
				URL:     url,
				Date:    date,
				Display: true,
				Venue:   venue,
			})
		}
	})

	c.Visit(SELLWOOD_CRAWL_URL)

	return kidsEvents
}

func getZooData() []KidsEvent {
	kidsEvents := []KidsEvent{}

	c := colly.NewCollector(
		colly.AllowedDomains(ZOO_DOMAIN),
	)

	c.OnHTML(".node-event", func(e *colly.HTMLElement) {
		titleText := e.ChildText(".node-title")
		url := ZOO_LINK_PREFIX + e.ChildAttr(".node-title a", "href")
		date := e.ChildAttr(".date-display-start", "content")
		venue := "Oregon Zoo"

		if titleText != "" {
			kidsEvents = append(kidsEvents, KidsEvent{
				Title:   titleText,
				URL:     url,
				Date:    date,
				Display: true,
				Venue:   venue,
			})
		}
	})

	c.Visit(ZOO_CRAWL_URL)

	return kidsEvents
}
