package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var coldStart = true
var initTime = time.Now()

func handler(ctx context.Context, event events.SNSEvent) error {
	start := time.Now()

	// Print full SNS event
	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("SNS Event:\n%s\n", string(eventJSON))

	// Print message(s)
	for _, record := range event.Records {
		fmt.Printf("SNS Message: %s\n", record.SNS.Message)
	}

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
