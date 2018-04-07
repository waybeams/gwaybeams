package display

type fakeWindow struct {
	ApplicationComponent
	frameRate int
}

func (w *fakeWindow) Init() {
}

func (w *fakeWindow) GetFrameRate() int {
	return 12
}

func (w *fakeWindow) PollEvents() []Event {
	return nil
}

func NewFakeWindow() Window {
	return &fakeWindow{}
}
