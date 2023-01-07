package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

type KidsEvent struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Date  string `json:"date"`
	Venue string `json:"venue"`
}

const (
	DOMAIN   = "www.portland5.com"
	FULL_URL = "https://" + DOMAIN
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	kidsEvents := []KidsEvent{}

	c := colly.NewCollector(
		colly.AllowedDomains(DOMAIN),
	)

	c.OnHTML(".views-row", func(e *colly.HTMLElement) {
		titleText := e.ChildText(".views-field-title .field-content a")
		url := FULL_URL + e.ChildAttr(".views-field-title .field-content a", "href")
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

	c.Visit("https://www.portland5.com/event-types/kids")

	b, err := json.Marshal(kidsEvents)
	if err != nil {
		log.Fatal(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/json"},
		Body:            string(b),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
