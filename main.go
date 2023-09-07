package main

import (
	"log"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	log.Println("HelloWorld from Golang Lambda")

	return nil
}

func main() {
	lambda.Start(handler)
}
