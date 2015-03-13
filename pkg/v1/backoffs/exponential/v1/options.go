package exponential

import (
	"time"
)

func WithLimit(value uint64) Option {
	return func(b *Backoff) {
		b.limit = value
	}
}

func WithFactor(value float64) Option {
	return func(b *Backoff) {
		b.factor = value
	}
}

func WithMaximum(value time.Duration) Option {
	return func(b *Backoff) {
		b.maximum = value
	}
}

func WithMinimum(value time.Duration) Option {
	return func(b *Backoff) {
		b.minimum = value
	}
}
