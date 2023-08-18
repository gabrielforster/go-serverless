package main

import (
  "fmt"
  "context"

  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
}

type Response struct {
  Message string `json:"message"`
}

func handler(ctx context.Context, event Event) (Response, error) {
  return Response{
    Message: fmt.Sprintf("Hello from Go running on a lambda!"),
  }, nil
}

func main () {
  lambda.Start(handler)
}
