package main

import (
	"context"
	"encoding/json"
	"log"

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
	kidsEvents, err := getDataFromDB()
	if err != nil {
		log.Fatal(err)
	}

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
