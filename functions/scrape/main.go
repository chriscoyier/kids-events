package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type KidsEvent struct {
	ID      string `json:"id"`
	Title   string `json:"title"` // e.g. "Storytime at the Zoo"
	URL     string `json:"url"`
	Date    string `json:"date"`
	Venue   string `json:"venue"`
	Display bool   `json:"display"` // Show in UI or not
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	kidsEvents := getPortland5Data()
	// kidsEvents = append(kidsEvents, getSellwoodData()...)
	// kidsEvents = append(kidsEvents, getZooData()...)

	saveToDB(kidsEvents)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "Scraping complete",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
