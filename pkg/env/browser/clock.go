package browser

type canvasClock struct {
}

func NewClock() *canvasClock {
	return &canvasClock{}
}

/*
func (w *window) OnFrame(handler func() bool, fps int, optClocks ...clock.Clock) {
	animFrame := w.browserWindow.Get("requestAnimationFrame")
	var wrapped func()

	wrapped = func() {
		handler()
		animFrame.Invoke(wrapped)
	}
	animFrame.Invoke(wrapped)
}
*/
