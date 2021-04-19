package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "fmt"
    "time"
)

func background() {
    time.Sleep(time.Second * 10)
    fmt.Println("OK")
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    fmt.Println("BEGIN")
    go background()
    return &events.APIGatewayProxyResponse {
        StatusCode: 200,
        Body: "Hello, World!",
    }, nil
}

func main() {
    lambda.Start(handler)
}
