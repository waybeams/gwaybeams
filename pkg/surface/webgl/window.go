package webgl

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
)

const DefaultFrameRate = 60
const DefaultHeight = 600
const DefaultTitle = "Default Title"
const DefaultWidth = 800

type window struct {
	events.EmitterBase

	frameRate  int
	height     float64
	pixelRatio float64
	title      string
	width      float64
}

func (w *window) BeginFrame() {
}

func (w *window) EndFrame() {
}

func (w *window) Close() {
}

func (w *window) FrameRate() int {
	return w.frameRate
}

func (w *window) GetCursorPos() (x, y float64) {
	return 0.0, 0.0
}

func (w *window) Init() {
}

func (w *window) OnResize(handler events.EventHandler) events.Unsubscriber {
	return nil
}

func (w *window) PixelRatio() float64 {
	return 1
}

func (w *window) PollEvents() {
}

func (w *window) ShouldClose() bool {
	return false
}

func (w *window) SetWidth(width float64) {
	w.width = width
}

func (w *window) SetHeight(height float64) {
	w.height = height
}

func (w *window) Width() float64 {
	return w.width
}

func (w *window) Height() float64 {
	return w.height
}

func (w *window) SetTitle(title string) {
	w.title = title
}

func (win *window) Title() string {
	return win.title
}

func (win *window) UpdateInput(root spec.ReadWriter) {
	panic("webgl.Window UpdateInput not implemented")
}

func NewWindow(options ...WindowOption) *window {
	defaults := []WindowOption{
		Width(DefaultWidth),
		Height(DefaultHeight),
		Title(DefaultTitle),
		FrameRate(DefaultFrameRate),
	}

	w := &window{}
	options = append(defaults, options...)
	for _, option := range options {
		option(w)
	}
	return w
}
