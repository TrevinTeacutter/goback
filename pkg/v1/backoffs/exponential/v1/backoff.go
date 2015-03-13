package exponential

import (
	"math"
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1"
)

const MaxAttemptsExceededError = basic("maximum of attempts exceeded")

var _ backoff.Backoff = (*Backoff)(nil)

type Option func(*Backoff)

type Backoff struct {
	// state
	attempts uint64
	next     time.Duration

	// configuration
	limit   uint64
	factor  float64
	minimum time.Duration
	maximum time.Duration
}

func New(options ...Option) (*Backoff, error) {
	b := &Backoff{}

	for _, option := range options {
		option(b)
	}

	return b, nil
}

func (b *Backoff) NextAttempt() (time.Duration, error) {
	if b.limit > 0 && b.attempts >= b.limit {
		return 0, MaxAttemptsExceededError
	}

	switch {
	case b.next >= b.maximum:
		b.next = b.maximum
	default:
		b.next = b.GetNextDuration()
	}

	if b.attempts < math.MaxUint64 {
		b.attempts++
	}

	return b.next, nil
}

func (b *Backoff) Reset() {
	b.attempts = 0
	b.next = 0
}

func (b *Backoff) GetNextDuration() time.Duration {
	duration := time.Duration(float64(b.minimum) * math.Pow(b.factor, float64(b.attempts)))

	if duration > b.maximum {
		return b.maximum
	}

	return duration
}
