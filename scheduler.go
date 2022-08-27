package main

import (
	"context"
	"fmt"
	"time"
)

// Schedule calls function `f` with a period `p` offsetted by `o`.
func Schedule(ctx context.Context, p time.Duration, o time.Duration, f func(time.Time)) {
	// Position the first execution
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	firstC := time.After(first.Sub(time.Now()))

	// Receiving from a nil channel blocks forever
	t := &time.Ticker{C: nil}

	for {
		select {
		case v := <-firstC:
			// The ticker has to be started before f as it can take some time to finish
			t = time.NewTicker(p)
			f(v)
		case v := <-t.C:
			f(v)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}

}

func main() {
	ctx := context.Background() //  .TODO()
	fmt.Println("Let's start:", time.Now())
	current := time.Now()
	hr := current.Hour()
	fmt.Printf("The stated hour "+
		"within the day is: %v\n", hr)
	// Schedule(ctx, time.Minute*2, time.Minute, fn()) Run every 2 minutes, starting 1 minute after the first run of this code
	// if the code started 52:41.26, then first run will be at 53:00 followed by another run at 55:00, and so on
	Schedule(ctx, time.Hour, time.Hour, func(t time.Time) {

		fmt.Println("Hi, it is:", time.Now())
	})
}
