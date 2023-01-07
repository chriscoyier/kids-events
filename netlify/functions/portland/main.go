package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	bodyString := ""

	fmt.Println("This message will show up in the CLI console.")

	c := colly.NewCollector(
		colly.AllowedDomains("www.portland5.com"),
	)

	c.OnHTML(".views-row", func(e *colly.HTMLElement) {
		bodyString += e.ChildText(".views-field-title .field-content a")
	})

	// c.OnError(func(_ *colly.Response, err error) {
	// 	fmt.Println("Something went wrong:", err)
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited", r.Request.URL)
	// })

	c.Visit("https://www.portland5.com/event-types/kids")

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/json"},
		Body:            bodyString,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
