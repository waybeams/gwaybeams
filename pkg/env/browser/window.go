package browser

import (
	"github.com/gopherjs/gopherjs/js"
	dom "github.com/oskca/gopherjs-dom"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
)

const DefaultFrameRate = 60
const DefaultHeight = 600
const DefaultTitle = "Default Title"
const DefaultWidth = 800

type window struct {
	events.EmitterBase

	browserWindow        *js.Object
	wrappedBrowserWindow *dom.Win
	frameRate            int
	height               float64
	pixelRatio           float64
	title                string
	titleChanged         bool
	width                float64
}

func (w *window) BeginFrame() {
	if w.titleChanged {
		w.browserWindow.Get("document").Set("title", w.title)
		w.titleChanged = false
	}
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
	w.wrappedBrowserWindow = dom.WrapWindow(w.browserWindow)
}

func (w *window) OnResize(handler events.EventHandler) events.Unsubscriber {
	bWin := w.browserWindow
	w.wrappedBrowserWindow.AddEventListener("resize", func(e *dom.Event) {
		// NOTE(lbayes): I'm getting zeros from the wrapped window for height/width.
		// not sure why, but pretty sure I'm doing something weird here.
		w.width = bWin.Get("innerWidth").Float()
		w.height = bWin.Get("innerHeight").Float()
	}, false)
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
	if title != w.title {
		w.titleChanged = true
	}
	w.title = title
}

func (w *window) Title() string {
	return w.title
}

func (w *window) UpdateInput(root spec.ReadWriter) {
	//panic("canvas.Window UpdateInput not implemented")
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
