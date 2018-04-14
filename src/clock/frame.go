package clock

import (
	"time"
)

type FrameHandler func()
type FrameUnsubscriber func()

func clockFromOptions(optClocks ...Clock) Clock {
	switch len(optClocks) {
	case 0:
		return New()
	case 1:
		return optClocks[0]
	default:
		panic("Only zero or one optional clocks are supported")
	}
}

func msPerFrame(fps int) time.Duration {
	return (time.Second / time.Duration(fps))
}

// OnFrame calls the provided handler at roughly the provided frames per second.
// The sleep duration will be offset by however long the last handler call blocked.
// This will give us an approximation of the desired frame rate and will drop
// frames for however long handlers take to execute.
// This call also accepts zero or one optional clock instances. This makes testing
// possible. If no clock is provided, the default real clock is used.
func OnFrame(handler FrameHandler, fps int, optClocks ...Clock) {
	clock := clockFromOptions(optClocks...)
	perFrame := msPerFrame(fps)

	for {
		startTime := clock.Now()
		handler()
		waitDuration := clock.Since(startTime) * time.Millisecond
		clock.Sleep(perFrame - waitDuration)
	}
}
