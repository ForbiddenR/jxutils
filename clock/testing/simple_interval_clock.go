package testing

import (
	"time"

	"github.com/ForbiddenR/jxutils/clock"
)

var (
	_ = clock.PassiveClock(&SimpleIntervalClock{})
)

// SimpleIntervalClock implements clock.PassiveClock, but each invocation of Now steps the clock forward the specified duration
type SimpleIntervalClock struct {
	Time     time.Time
	Duration time.Duration
}

// Now returns i's time.
func (i *SimpleIntervalClock) Now() time.Time {
	i.Time = i.Time.Add(i.Duration)
	return i.Time
}

// Since returns time since the time in i.
func (i *SimpleIntervalClock) Since(ts time.Time) time.Duration {
	return i.Time.Sub(ts)
}