package main

import (
	"github.com/gocolly/colly"
)

const (
	PORTLAND5_DOMAIN      = "www.portland5.com"
	PORTLAND5_LINK_PREFIX = "https://" + PORTLAND5_DOMAIN
	PORTLAND5_CRAWL_URL   = "https://www.portland5.com/event-types/kids"

	SELLWOOD_DOMAIN      = "www.sellwoodcommunityhouse.org"
	SELLWOOD_LINK_PREFIX = "https://" + SELLWOOD_DOMAIN
	SELLWOOD_CRAWL_URL   = "https://www.sellwoodcommunityhouse.org/community-events-1"
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
				Title: titleText,
				URL:   url,
				Date:  date,
				Venue: venue,
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
		date := ""

		venue := "Sellwood Community House"

		if titleText != "" {
			kidsEvents = append(kidsEvents, KidsEvent{
				Title: titleText,
				URL:   url,
				Date:  date,
				Venue: venue,
			})
		}
	})

	c.Visit(SELLWOOD_CRAWL_URL)

	return kidsEvents
}
