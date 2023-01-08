package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type KidsEvent struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Date  string `json:"date"`
	Venue string `json:"venue"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	kidsEvents := getPortland5Data()
	kidsEvents = append(kidsEvents, getSellwoodData()...)
	kidsEvents = append(kidsEvents, getZooData()...)

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
