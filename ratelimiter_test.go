package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/charles-haynes/ratelimiter"
)

func TestEmptyLimiterDoesNotLimit(t *testing.T) {
	l := ratelimiter.New(0, 50*time.Millisecond)
	i := 0
	for t := time.Now(); time.Since(t) < 50*time.Millisecond; {
		l.Limit()
		i++
	}
	if i <= 1000 {
		t.Errorf("expected > 1000 iterations, got %d", i)
	}
}

func TestZeroTimeLimiterDoesNotLimit(t *testing.T) {
	l := ratelimiter.New(1, 0)
	i := 0
	for t := time.Now(); time.Since(t) < 50*time.Millisecond; {
		l.Limit()
		i++
	}
	if i <= 1000 {
		t.Errorf("expected > 1000 iterations, got %d", i)
	}
}

// check that a uniform limiter works
func TestLimiterLimits(t *testing.T) {
	l := ratelimiter.New(1, 1*time.Millisecond)
	i := 0
	for t := time.Now(); time.Since(t) < 100*time.Millisecond; {
		l.Limit()
		i++
	}
	if 90 > i || i > 101 {
		t.Errorf("expected about 100 iterations, got %d", i)
	}
}

// make sure a bursty limiter averages out properly over multiple
// periods
func TestLimiterLimitsBurst(t *testing.T) {
	l := ratelimiter.New(10, 10*time.Millisecond)
	i := 0
	for t := time.Now(); time.Since(t) < 100*time.Millisecond; {
		l.Limit()
		i++
	}
	if 90 > i || i > 101 {
		t.Errorf("expected about 100 iterations, got %d", i)
	}
}

// make sure a burst of 100 calls is allowed in a short amount of time
func TestLimiterLimitsBigBurst(t *testing.T) {
	l := ratelimiter.New(100, 100*time.Millisecond)
	i := 0
	for t := time.Now(); time.Since(t) < 10*time.Millisecond; {
		l.Limit()
		i++
	}
	if 90 > i || i > 101 {
		t.Errorf("expected about 100 iterations, got %d", i)
	}
}
