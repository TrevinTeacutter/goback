package backoff

import (
	"context"
	"time"
)

type Backoff interface {
	NextAttempt() (time.Duration, error)
	Reset()
}

func Wait(ctx context.Context, b Backoff) error {
	next, err := b.NextAttempt()
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(next):
		return nil
	}
}
