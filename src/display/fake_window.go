package display

type fakeWindow struct {
	width     float64
	height    float64
	frameRate int
	title     string
}

func (w *fakeWindow) Init() {
}

func (w *fakeWindow) GetFrameRate() int {
	return 12
}

func (w *fakeWindow) GetHeight() float64 {
	return w.height
}

func (w *fakeWindow) GetTitle() string {
	return w.title
}

func (w *fakeWindow) GetWidth() float64 {
	return w.width
}

func (w *fakeWindow) Height(h float64) {
	w.height = h
}

func (w *fakeWindow) PollEvents() []Event {
	return nil
}

func (w *fakeWindow) Title(str string) {
	w.title = str
}

func (w *fakeWindow) Width(width float64) {
	w.width = width
}

func NewFakeWindow() Window {
	return &fakeWindow{}
}
