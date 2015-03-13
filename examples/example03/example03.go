package main

import (
	"context"
	"fmt"
	"log"
	"time"

	backoff "github.com/TrevinTeacutter/goback/pkg/v1"
	"github.com/TrevinTeacutter/goback/pkg/v1/backoffs/exponential/v1"
)

// Creates a function that will fail 6 times before connecting
func retryGenerator() func(chan bool) {
	var failedAttempts = 6
	return func(done chan bool) {
		if failedAttempts > 0 {
			failedAttempts--
			return
		}
		done <- true
	}
}

func connect(ctx context.Context, retry func(chan bool), b backoff.Backoff) {
	done := make(chan bool, 1)
	for {
		retry(done)
		select {
		case err := <-backoff.After(ctx, b):
			if err != nil {
				log.Fatalf("Error connecting: %v", err)
			}
			log.Printf("Problem connecting")
			continue
		case <-done:
			//conected
			b.Reset()
			return
		}
	}

}

func main() {
	ctx := context.Background()
	retry := retryGenerator()

	b, err := exponential.New(
		exponential.WithMinimum(100*time.Millisecond),
		exponential.WithMaximum(60*time.Second),
		exponential.WithFactor(2),
		exponential.WithLimit(4),
	)
	if err != nil {
		panic(err)
	}

	connect(ctx, retry, b)
	// Duplicates the time each time from a minimum of 100ms to a maximum of 1 min.
	fmt.Println("Connected")
}
