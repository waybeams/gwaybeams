package display

import (
	"time"
)

// ApplicationComponent belongs at the root of any component tree that will
// manage change over time.
type ApplicationComponent struct {
	Component
	frameRate int
}

func (a *ApplicationComponent) GetFrameRate() int {
	if a.frameRate == 0 {
		return DefaultFrameRate
	}
	return a.frameRate
}

func (a *ApplicationComponent) GetFrameStart() time.Time {
	return time.Now()
}

func (a *ApplicationComponent) WaitForFrame(startTime time.Time) {
	// Wait for whatever amount of time remains between how long we just spent,
	// and when the next frame (at fps) should be.
	waitDuration := (time.Second / time.Duration(a.GetFrameRate())) - time.Since(startTime)
	// NOTE: Looping stops when mouse is pressed on window resizer (on macOS, but not i3wm/Ubuntu Linux)
	time.Sleep(waitDuration)
}

func NewApplication() Displayable {
	return &ApplicationComponent{}
}

var Application = NewComponentFactory("Application", NewApplication)
