package main

import (
	"log"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	log.Println("Hello World from Golang Lambda")

	return "funcionou", nil
}

func main() {
	lambda.Start(handler)
}
