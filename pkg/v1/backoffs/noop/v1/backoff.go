package noop

import (
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1"
)

var _ backoff.Backoff = (*Backoff)(nil)

type Backoff struct{}

func (b *Backoff) NextAttempt() (time.Duration, error) {
	return 0, nil
}

func (b *Backoff) Reset() {}
