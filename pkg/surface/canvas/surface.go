package canvas

import (
	"github.com/gopherjs/gopherjs/js"
	jsCanvas "github.com/oskca/gopherjs-canvas"
	"github.com/waybeams/waybeams/pkg/spec"
)

type Surface struct {
	context     *jsCanvas.Canvas
	pageContext *js.Object
	window      *js.Object
	flags       []SurfaceOption
	width       float64
	height      float64
	fonts       map[string]spec.Font
}

func (s *Surface) Init() {
	s.context = jsCanvas.New(s.pageContext)
	s.window = js.Global.Get("window")
	/*
		pageContext.Set("width", win.Get("innerWidth"))
		pageContext.Set("height", win.Get("innerHeight"))
	*/

	/*
		context, err := nanovgo.NewContext(s.Flags())
		if err != nil {
			panic(err)
		}

		s.context = context
	*/
}

func (s *Surface) Close() {
	/*
		if s.context != nil {
			s.context.Delete()
		}
	*/
}

func (s *Surface) BeginFrame(w, h float64) {
	if w != s.width {
		s.pageContext.Set("width", w)
		s.width = w
	}
	if h != s.height {
		s.pageContext.Set("height", h)
		s.height = h
	}
	/*
		s.CreateFonts()

		// ratio := float32(w / h)
		s.context.BeginFrame(int(w), int(h), 1)
	*/
}

func (s *Surface) EndFrame() {
	/*
		s.context.EndFrame()
	*/
}

func (s *Surface) getFonts() map[string]spec.Font {
	if s.fonts == nil {
		s.fonts = make(map[string]spec.Font)
	}
	return s.fonts
}

func (s *Surface) AddFont(name string, path string) {
	/*
		fonts := s.getFonts()
		if fonts[name] == nil {
			fonts[name] = spec.NewFont(name, path)
		}
	*/
}

func (s *Surface) CreateFonts() {
	/*
		for _, font := range s.getFonts() {
			if !font.IsCreated() {
				s.CreateFont(font.Name(), font.Path())
				font.OnCreated()
			}
		}
	*/
}

func (s *Surface) Font(name string) spec.Font {
	return s.getFonts()[name]
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
	/*
		s.context.CreateFont(name, path)
	*/
}

func (s *Surface) MoveTo(x float64, y float64) {
	// s.context.MoveTo(float32(x), float32(y))
}

func (s *Surface) SetFillColor(color uint) {
	// r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	// s.context.SetFillColor(nanovgo.Color{r, g, b, a})
}

func (s *Surface) SetStrokeColor(color uint) {
	// r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	// s.context.SetStrokeColor(nanovgo.Color{r, g, b, a})
}

func (s *Surface) SetStrokeWidth(width float64) {
	// s.context.SetStrokeWidth(float32(width))
}

func (s *Surface) Stroke() {
	// s.context.Stroke()
}

func (s *Surface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	// TODO(lbayes): Update external Surface to include direction and facilitate for Cairo
	// s.context.Arc(float32(xc), float32(yc), float32(radius), float32(angle1), float32(angle2), nanovgo.Clockwise)
}

func (s *Surface) BeginPath() {
	// s.context.BeginPath()
}

func (s *Surface) DebugDumpPathCache() {
	// s.context.DebugDumpPathCache()
}

func (s *Surface) Fill() {
	// s.context.Fill()
}

func (s *Surface) Rect(x, y, width, height float64) {
	// s.context.Rect(float32(x), float32(y), float32(width), float32(height))
}

func (s *Surface) RoundedRect(x, y, width, height, radius float64) {
	// s.context.RoundedRect(float32(x), float32(y), float32(width), float32(height), float32(radius))
}

func (s *Surface) SetFontSize(size float64) {
	// s.context.SetFontSize(float32(size))
}

func (s *Surface) SetFontFace(face string) {
	// s.context.SetFontFace(face)
}

func (s *Surface) Text(x float64, y float64, text string) {
	// TODO(lbayes): Add validation that ensures required calls have been made before calling this function (e.g., SetFontFace)
	// s.context.Text(float32(x), float32(y), text)
}

type SurfaceOption func(s *Surface)

/*
func Width(width float64) SurfaceOption {
	return func(s *Surface) {
		// s.SetWidth(width)
	}
}

func Height(height float64) SurfaceOption {
	return func(s *Surface) {
		// s.SetHeight(height)
	}
}

func AntiAlias() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.AntiAlias)
	}
}

func StencilStrokes() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.StencilStrokes)
	}
}

func Debug() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.Debug)
	}
}
*/

func PageContext(context *js.Object) SurfaceOption {
	return func(s *Surface) {
		s.pageContext = context
	}
}

func Font(name, path string) SurfaceOption {
	return func(s *Surface) {
		s.AddFont(name, path)
	}
}

func NewSurface(options ...SurfaceOption) *Surface {
	s := &Surface{}

	for _, option := range options {
		option(s)
	}

	if s.pageContext == nil {
		panic("Surface(PageContext(...)) is required")
	}
	return s
}
