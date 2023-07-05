package testing

import (
	"testing"
	"time"
)

func TestSimpleIntervalClockNow(t *testing.T) {
	cases := []struct {
		duration time.Duration
	}{
		{duration: 10 * time.Millisecond},
		{duration: 0 * time.Millisecond},
		{duration: -10 * time.Millisecond},
	}

	startTime := time.Now()
	for _, c := range cases {
		expectedTime := startTime.Add(c.duration)
		sic := &SimpleIntervalClock{
			Time:     startTime,
			Duration: c.duration,
		}
		actualTime := sic.Now()
		if !expectedTime.Equal(actualTime) {
			t.Errorf("expected %#v, got %#v", expectedTime, actualTime)
		}
	}
}

func TestSimpleIntervalClockSince(t *testing.T) {
	cases := []struct {
		delta time.Duration
	}{
		{delta: 10 * time.Millisecond},
		{delta: 0 * time.Millisecond},
		{delta: -10 * time.Millisecond},
	}

	startTime := time.Now()
	duration := time.Millisecond
	sic := &SimpleIntervalClock{
		Time:     startTime,
		Duration: duration,
	}

	for _, c := range cases {
		// Try and add compute a "since" time by
		// Add()ing a -c.delta.
		timeSinceDelta := startTime.Add(-c.delta)
		expectedDelta := sic.Since(timeSinceDelta)
		if expectedDelta != c.delta {
			t.Errorf("expected %#v, got %#v", expectedDelta, c.delta)
		}
	}
}