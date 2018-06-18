package browser

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	jsCanvas "github.com/oskca/gopherjs-canvas"
	"github.com/waybeams/waybeams/pkg/spec"
)

type metricsStoreType map[string]*js.Object

const DefaultFontSize = 10

type Font struct {
	context      *jsCanvas.Context2D
	surface      *Surface
	name         string
	path         string
	size         float64
	created      bool
	stash        interface{}
	metricsStore metricsStoreType
}

func (f *Font) getMetricsStore() metricsStoreType {
	if f.metricsStore == nil {
		f.metricsStore = make(metricsStoreType)
	}
	return f.metricsStore
}

func (f *Font) getMetricsFor(value string, name string, size float64) *js.Object {
	f.surface.SetFontSize(size)
	f.surface.SetFontFace(name)

	return f.context.Get("measureText").Invoke(value)
}

func (f *Font) Name() string {
	return f.name
}

func (f *Font) Path() string {
	return f.path
}

func (f *Font) Size() float64 {
	if f.size == 0.0 {
		f.size = DefaultFontSize
	}
	return f.size
}

func (f *Font) SetSurface(surface *Surface) {
	f.surface = surface
}

func (f *Font) SetContext(context *jsCanvas.Context2D) {
	f.context = context
}

func (f *Font) OnCreated() {
	f.created = true
}

func (f *Font) IsCreated() bool {
	return f.created
}

func (f *Font) getStash() interface{} {
	if f.stash == nil {
		// f.stash = fsm.New(nvgInitFontImageSize, nvgInitFontImageSize)
		// result := f.stash.AddFont(f.name, f.path)
		// if result == -1 {
		// msg := "Unable to add font, likely bad path: " + f.path
		// panic(msg)
		// }
	}
	return f.stash
}

func (f *Font) SetSize(size float64) {
	f.size = size
}

func (f *Font) SetAlign(align spec.Alignment) {
	// f.getStash().SetAlign(fsa)
}

func (f *Font) VerticalMetrics() (ascender, descender, lineHeight float64) {
	// f.context.VerticalMetrics
	// stash := f.getStash()
	// return stash.VerticalMetrics()
	return 0.0, 0.0, f.Size()
}

func (f *Font) Bounds(value string) (width float64, bounds []float64) {
	name := f.Name()
	size := f.Size()
	metrics := f.getMetricsFor(value, name, size)

	// Info on font/actual boundingBox ascent/descent here:
	// https://stackoverflow.com/questions/46949891/html5-canvas-fontboundingboxascent-vs-actualboundingboxascent

	w := metrics.Get("width").Float()
	fmt.Println("METRICS:", w, size)
	// metrics.fontBoundingBoxAscent + metrics.fontBoundingBoxDescent;
	ascent := metrics.Get("fontBoundingBoxAscent")
	keys := js.Keys(metrics)
	fmt.Println("KEYS:", keys)
	// descent := metrics.Get("fontBoundingBoxDescent")
	descent := 0.0
	// fmt.Println("metrics>>>>> :", w)
	fmt.Println("ascent:", ascent)
	fmt.Println("descent:", descent)

	// return stash.TextBounds(0, 0, value)
	return w, []float64{0, 0, w, size}
}

func NewFont(name string, path string) *Font {
	return &Font{
		name: name,
		path: path,
	}
}
