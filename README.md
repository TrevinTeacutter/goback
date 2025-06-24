# goback

Goback implements a simple exponential backoff.

An exponential backoff approach is typically used when treating with potentially faulty/slow systems. If a system fails quick retries may exacerbate the system specially when the system is dealing with several clients. In this case a backoff provides the faulty system enough room to recover.

## How to use
```go
func main() {
    b, err := exponential.New(
        exponential.WithMinimum(100*time.Millisecond),
        exponential.WithMaximum(60*time.Second),
        exponential.WithFactor(2), 
    )
    if err != nil {
        panic(err)
    }

    backoff.Wait(b)               // sleeps 100ms
    backoff.Wait(b)               // sleeps 200ms
    backoff.Wait(b)               // sleeps 400ms
    fmt.Println(b.NextDuration()) // prints 800ms
    b.Reset()                     // resets the backoff
    backoff.Wait(b)               // sleeps 100ms
}
```

Further examples can be found in the examples folder.

## Strategies
At the moment there are four backoff implementations.

### No-Op Backoff
This is a backoff implementation that will always return zero values for `NextDuration`.

### Constant Backoff
This is a backoff implementation that will always return the given value for `NextDuration`.

### Exponential Backoff
It starts with a minumum duration and multiplies it by the factor until the maximum waiting time is reached. In that case it will return `Maximum`.

The optional `Limit` will limit the maximum number of retries and will return an error when is exceeded.

### Jitter Backoff
The Jitter strategy leverages a provided backoff implementation but adds a light randomisation to minimise collisions between contending clients.

### Extensibility
By creating structs that implement the methods of the `Backoff` interface you will be able to use them as a backoff strategy.

A naive example of this is:
```go
type NaiveBackoff struct{}

func (b *NaiveBackoff) NextAttempt() (time.Duration, error) { return time.Second, nil }
func (b *NaiveBackoff) Reset() { }
```
This will return always a 1s duration.

## Credits
This package is based on https://github.com/carlescere/goback with a larger focus on implementations having more of a single purpose and leveraging composition to layer desired behavior of the resulting structure.
