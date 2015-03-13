package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1"
	"github.com/TrevinTeacutter/goback/pkg/v1/backoffs/exponential/v1"
)

// Creates a function that will fail 6 times before connecting
func faultyTCPGenerator() func(string, string, string) (net.Listener, error) {
	var failedAttempts = 6
	
	return func(protocol, address, port string) (net.Listener, error) {
		if failedAttempts > 0 {
			failedAttempts--

			return nil, fmt.Errorf("haha") // :)
		}

		l, err := net.Listen(protocol, fmt.Sprintf("%s:%s", address, port))

		return l, err
	}
}

func main() {
	ctx := context.Background()
	tcpConnect := faultyTCPGenerator()

	// Duplicates the time each time from a minimum of 100ms to a maximum of 1 min.
	b, err := exponential.New(
		exponential.WithMinimum(100*time.Millisecond),
		exponential.WithMaximum(60*time.Second),
		exponential.WithFactor(2),
	)
	if err != nil {
		panic(err)
	}

	for {
		l, inner := tcpConnect("tcp", "localhost", "5000")
		if inner != nil { // fail to connect
			log.Printf("Error connecting: %v", err)

			_ = backoff.Wait(ctx, b) // Exponential backoff

			continue
		}

		defer l.Close()

		// connected
		log.Printf("Connected!")
		b.Reset() // Reset number of attempts. Brings backoff time to the minimum

		break // Here be dragons
	}
}
