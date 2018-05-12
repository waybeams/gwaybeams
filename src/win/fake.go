package win

type FakeWindow struct {
	width      float64
	height     float64
	pixelRatio float64
	frameRate  int
}

func (f *FakeWindow) FrameRate() int {
	return f.frameRate
}

func (f *FakeWindow) SetWidth(width float64) {
	f.width = width
}

func (f *FakeWindow) Width() float64 {
	return f.width
}

func (f *FakeWindow) SetHeight(height float64) {
	f.height = height
}

func (f *FakeWindow) Height() float64 {
	return f.height
}

func (f *FakeWindow) BeginFrame() {
}

func (f *FakeWindow) EndFrame() {
}

func (f *FakeWindow) Close() {
}

func (f *FakeWindow) PixelRatio() float64 {
	return f.pixelRatio
}

func (f *FakeWindow) ShouldClose() bool {
	return false
}

func NewFake() *FakeWindow {
	return &FakeWindow{}
}
