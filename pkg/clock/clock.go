package clock

import (
	"time"

	benclock "github.com/benbjohnson/clock"
)

type FrameHandler func() bool

// Clock represents an interface to the functions in the standard library time
// package. Two implementations are available in the clock package. The first
// is a real-time clock which simply wraps the time package's functions. The
// second is a mock clock which will only make forward progress when
// programmatically adjusted.
type Clock interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *benclock.Timer
	Now() time.Time
	OnFrame(handler FrameHandler, fps int)
	Since(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Ticker(d time.Duration) *benclock.Ticker
	Timer(d time.Duration) *benclock.Timer
}

func msPerFrame(fps int) time.Duration {
	return (time.Second / time.Duration(fps))
}

func onFrame(c Clock, handler FrameHandler, fps int) {
	perFrame := msPerFrame(fps)
	for {
		startTime := c.Now()
		shouldExit := handler()
		if shouldExit {
			return
		}
		workDuration := c.Since(startTime)
		c.Sleep(perFrame - workDuration)
	}
}

type clock struct {
	grossDelegate benclock.Clock
}

func (c *clock) After(d time.Duration) <-chan time.Time {
	return c.grossDelegate.After(d)
}
func (c *clock) AfterFunc(d time.Duration, f func()) *benclock.Timer {
	return c.grossDelegate.AfterFunc(d, f)
}

func (c *clock) Now() time.Time {
	return c.grossDelegate.Now()

}

func (c *clock) OnFrame(handler FrameHandler, fps int) {
	onFrame(c, handler, fps)
}

func (c *clock) Since(t time.Time) time.Duration {
	return c.grossDelegate.Since(t)
}

func (c *clock) Sleep(d time.Duration) {
	c.grossDelegate.Sleep(d)
}

func (c *clock) Tick(d time.Duration) <-chan time.Time {
	return c.grossDelegate.Tick(d)
}

func (c *clock) Ticker(d time.Duration) *benclock.Ticker {
	return c.grossDelegate.Ticker(d)
}

func (c *clock) Timer(d time.Duration) *benclock.Timer {
	return c.grossDelegate.Timer(d)
}

// New returns an instance of a real-time clock.
func New() Clock {
	return &clock{
		grossDelegate: benclock.New(),
	}
}
