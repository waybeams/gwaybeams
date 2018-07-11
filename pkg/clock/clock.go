package clock

import (
	"time"

	benclock "github.com/benbjohnson/clock"
)

// Clock represents an interface to the functions in the standard library time
// package. Two implementations are available in the clock package. The first
// is a real-time clock which simply wraps the time package's functions. The
// second is a mock clock which will only make forward progress when
// programmatically adjusted.
type Clock interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *benclock.Timer
	Now() time.Time
	Since(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Ticker(d time.Duration) *benclock.Ticker
	Timer(d time.Duration) *benclock.Timer
}

// New returns an instance of a real-time clock.
func New() Clock {
	return benclock.New()
}

// FakeClock adds test-only features to the Clock interface.
type FakeClock interface {
	Clock

	Add(d time.Duration)
	Set(t time.Time)
}

// NewFake returns an instance of the fake clock.
func NewFake() *benclock.Mock {
	return benclock.NewMock()
}
