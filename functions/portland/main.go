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
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	kidsEvents := []KidsEvent{}

	c := colly.NewCollector(
		colly.AllowedDomains("www.portland5.com"),
	)

	c.OnHTML(".views-row", func(e *colly.HTMLElement) {
		kidsEvents = append(kidsEvents, KidsEvent{
			Title: e.ChildText(".views-field-title .field-content a"),
			URL:   e.ChildAttr(".views-field-title .field-content a", "href"),
			Date:  e.ChildAttr(".date-display-single", "content"),
		})
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
