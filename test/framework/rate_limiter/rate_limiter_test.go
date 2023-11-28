package rate_limiter

import (
	"golang.org/x/time/rate"
	"testing"
	"time"
)

//learn from k8s.io/client-go/util/flowcontrol/throttle_test.go

func TestMultithreadedThrottling(t *testing.T) {
	r := rate.NewLimiter(100, 1)
	// channel to collect 100 tokens
	taken := make(chan bool, 100)

	// Set up goroutines to hammer the throttler
	startCh := make(chan bool)
	endCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			// wait for the starting signal
			<-startCh
			for {
				// get a token
				allow := r.Allow()
				if allow {
					select {
					// try to add it to the taken channel
					case taken <- true:
						continue
					// if taken is full, notify and return
					default:
						endCh <- true
						return
					}
				}
			}
		}()
	}

	// record wall time
	startTime := time.Now()
	// take the initial capacity so all tokens are the result of refill
	r.Allow()
	// start the thundering herd
	close(startCh)
	// wait for the first signal that we collected 100 tokens
	<-endCh
	// record wall time
	endTime := time.Now()

	// tolerate a 1% clock change because these things happen
	if duration := endTime.Sub(startTime); duration < (time.Second * 99 / 100) {
		// We shouldn't be able to get 100 tokens out of the bucket in less than 1 second of wall clock time, no matter what
		t.Errorf("Expected it to take at least 1 second to get 100 tokens, took %v", duration)
	} else {
		t.Logf("Took %v to get 100 tokens", duration)
	}
}
