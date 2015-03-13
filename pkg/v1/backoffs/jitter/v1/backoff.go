package jitter

import (
	"math/rand"
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1"
)

var _ backoff.Backoff = (*Backoff)(nil)

type Backoff struct {
	// dependencies
	child backoff.Backoff

	// configuration
	jitter time.Duration
}

func New(child backoff.Backoff, jitter time.Duration) *Backoff {
	return &Backoff{
		child:  child,
		jitter: jitter,
	}
}

func (b *Backoff) NextAttempt() (time.Duration, error) {
	next, err := b.child.NextAttempt()
	if err != nil {
		return 0, err
	}

	return b.addJitter(next), nil
}

func (b *Backoff) Reset() {
	b.child.Reset()
}

func (b *Backoff) addJitter(next time.Duration) time.Duration {
	return time.Duration(rand.Float64()*float64(2*b.jitter) + float64(next-b.jitter))
}
