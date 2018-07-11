package clock

import (
	"time"

	benclock "github.com/benbjohnson/clock"
)

// FakeClock adds test-only features to the Clock interface.
type Fake interface {
	Clock

	Add(d time.Duration)
	Set(t time.Time)
}

type fake struct {
	benclock.Mock
}

func (c *fake) OnFrame(handler FrameHandler, fps int) {
	onFrame(c, handler, fps)
}

// NewFake returns an instance of the fake clock.
func NewFake() Fake {
	return &fake{}
}
