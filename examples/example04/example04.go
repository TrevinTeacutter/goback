package main

import (
	"fmt"
	"time"

	"github.com/TrevinTeacutter/goback/pkg/v1/backoffs/exponential/v1"
	"github.com/TrevinTeacutter/goback/pkg/v1/backoffs/jitter/v1"
)

func main() {
	inner, err := exponential.New(
		exponential.WithMinimum(100*time.Millisecond),
		exponential.WithMaximum(60*time.Second),
		exponential.WithFactor(2),
	)
	if err != nil {
		panic(err)
	}

	b := jitter.New(inner, time.Second*2)
	
	fmt.Println(b.NextAttempt())
	fmt.Println(b.NextAttempt())
	fmt.Println(b.NextAttempt())
	fmt.Println(b.NextAttempt())
	fmt.Println(b.NextAttempt())

}
