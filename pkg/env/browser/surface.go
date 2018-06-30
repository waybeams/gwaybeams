package browser

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"

	jsCanvas "github.com/oskca/gopherjs-canvas"
	"github.com/waybeams/waybeams/pkg/helpers"
)

type ExternalCanvas interface {
	Set(string, interface{})
	GetContext2D() *jsCanvas.Context2D
}

const Clockwise = false
const Anticlockwise = true

type Surface struct {
	context *jsCanvas.Context2D
	canvas  ExternalCanvas

	flags  []SurfaceOption
	width  float64
	height float64

	lastFontSize    int
	lastFontFace    string
	lastStrokeWidth int
	lastStrokeColor uint
}

func (s *Surface) Init() {
	s.context = s.canvas.GetContext2D()
}

func (s *Surface) Close() {
}

func (s *Surface) BeginFrame(w, h float64) {
	if w != s.width {
		s.canvas.Set("width", w)
		s.width = w
	}
	if h != s.height {
		s.canvas.Set("height", h)
		s.height = h
	}
}

func (s *Surface) EndFrame() {
}

func (s *Surface) SetWidth(width float64) {
	s.width = width
}

func (s *Surface) SetHeight(height float64) {
	s.height = height
}

func (s *Surface) Width() float64 {
	return s.width
}

func (s *Surface) Height() float64 {
	return s.height
}

func (s *Surface) CreateFont(name, path string) {
	panic("Not implemented")
}

func (s *Surface) MoveTo(x float64, y float64) {
	s.context.MoveTo(x, y)
}

func (s *Surface) SetFillColor(color uint) {
	s.context.FillStyle = helpers.UintToHexString(color)
}

func (s *Surface) SetStrokeColor(color uint) {
	s.context.StrokeStyle = helpers.UintToHexString(color)
}

func (s *Surface) SetStrokeWidth(width float64) {
	s.context.LineWidth = width
}

func (s *Surface) Stroke() {
	s.context.Stroke()
}

func (s *Surface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	s.context.Arc(xc, yc, radius, angle1, angle2, Clockwise)
}

func (s *Surface) BeginPath() {
	s.context.BeginPath()
}

func (s *Surface) DebugDumpPathCache() {
	panic("DebugDumpPathCache not available in HTML Canvas")
}

func (s *Surface) Fill() {
	s.context.Fill()
}

func (s *Surface) Rect(x, y, width, height float64) {
	s.context.Rect(x, y, width, height)
}

func (s *Surface) RoundedRect(x, y, width, height, radius float64) {
	// s.context.RoundedRect(x, y, width, height, radius)
}

func (s *Surface) AddFont(name string, path string) {
	// TODO(lbayes): Load font if path is URL?
	panic("Not Implemented")
}

func (s *Surface) SetFontSize(size float64) {
	s.lastFontSize = int(size)
}

func (s *Surface) SetFontFace(face string) {
	s.lastFontFace = face
}

func (s *Surface) TextBounds(face string, size float64, text string) (x, y, w, h float64) {
	// Info on font/actual boundingBox ascent/descent here:
	// https://stackoverflow.com/questions/46949891/html5-canvas-fontboundingboxascent-vs-actualboundingboxascent

	// w := metrics.Get("width").Float()
	// fmt.Println("METRICS:", w
	// metrics.fontBoundingBoxAscent + metrics.fontBoundingBoxDescent;
	// ascent := metrics.Get("fontBoundingBoxAscent")
	// keys := js.Keys(metrics)
	// fmt.Println("KEYS:", keys)
	// // descent := metrics.Get("fontBoundingBoxDescent")
	// descent := 0.0
	// // fmt.Println("metrics>>>>> :", w)
	// fmt.Println("ascent:", ascent)
	// fmt.Println("descent:", descent)
	// return stash.TextBounds(0, 0, value)

	// Fake values here:
	x = -0.5
	y = -18.0
	h = size
	w = (size * float64(len(text))) * 0.55
	return x, y, w, h
}

func (s *Surface) Text(x float64, y float64, text string) {
	// maxWidth required for canvas filltext, but not nanovgo?
	maxWidth := 10000000.0
	s.context.Font = strconv.Itoa(s.lastFontSize) + "px \"Open Sans\", sans-serif"
	// TODO(lbayes): Add validation that ensures required calls have been made before calling this function (e.g., SetFontFace)
	s.context.FillText(text, x, y, maxWidth)
}

// NewCanvasFromJsObject will wrap the provided GopherJs element with the
// Canvas wrapper provided by oska.
func NewCanvasFromJsObject(element *js.Object) *jsCanvas.Canvas {
	return jsCanvas.New(element)
}

// NewSurface will create a new Waybeams Surface for the gopherjs
// environment.
func NewSurface(canvas ExternalCanvas, options ...SurfaceOption) *Surface {
	s := &Surface{canvas: canvas}

	for _, option := range options {
		option(s)
	}

	if s.canvas == nil {
		panic("Surface(Canvas(...)) is required")
	}
	return s
}
