package constant

import (
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1"
)

var _ backoff.Backoff = (*Backoff)(nil)

type Backoff struct {
	constant time.Duration
}

func New(duration time.Duration) *Backoff {
	return &Backoff{constant: duration}
}

func (b *Backoff) NextAttempt() (time.Duration, error) {
	return b.constant, nil
}

func (b *Backoff) Reset() {}
