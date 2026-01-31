
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var coldStart = true
var initTime = time.Now()

func handler(ctx context.Context) error {
	start := time.Now()

	fmt.Printf("Language=Go\n")
	fmt.Printf("ColdStart=%v\n", coldStart)
	fmt.Printf("InitDurationMs=%d\n", time.Since(initTime).Milliseconds())

	coldStart = false

	fmt.Printf("ExecutionDurationMs=%d\n", time.Since(start).Milliseconds())
	return nil
}

func main() {
	lambda.Start(handler)
}
