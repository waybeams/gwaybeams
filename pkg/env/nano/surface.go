package nano

import (
    "fmt"
	"github.com/shibukawa/nanovgo"
	"github.com/waybeams/waybeams/pkg/helpers"
)

const fakePixelRatio = float32(1.0)

type Surface struct {
    context    *nanovgo.Context
    flags      []nanovgo.CreateFlags
    width      float64
    height     float64
    fonts      map[string]*Font
    pixelRatio float32
}

func (s *Surface) Init() {
	context, err := nanovgo.NewContext(s.Flags())
	if err != nil {
		panic(err)
	}
	s.context = context
    s.pixelRatio = 1
}

func (s *Surface) SetPixelRatio(ratio float32) {
    s.pixelRatio = ratio
}

func (s *Surface) SetScale(x, y float32) {
    fmt.Println("Nanovgo.Scale with w: and h: ", x, y)
	s.context.Scale(x, y)
}

func (s *Surface) Close() {
	if s.context != nil {
		s.context.Delete()
	}
}

func (s *Surface) BeginFrame() {
	s.CreateFonts()
    fmt.Println("Surface.BeginFrame with:", s.Width(), s.Height())
	s.context.BeginFrame(int(s.Width()), int(s.Height()), s.pixelRatio)
}

func (s *Surface) EndFrame() {
	s.context.EndFrame()
}

func (s *Surface) getFonts() map[string]*Font {
	if s.fonts == nil {
		s.fonts = make(map[string]*Font)
	}
	return s.fonts
}

func (s *Surface) AddFont(name string, path string) {
	fonts := s.getFonts()
	if fonts[name] == nil {
		fonts[name] = NewFont(name, path)
	}
}

func (s *Surface) CreateFonts() {
	for _, font := range s.getFonts() {
		if !font.IsCreated() {
			s.CreateFont(font.Name(), font.Path())
			font.OnCreated()
		}
	}
}

func (s *Surface) Font(name string) *Font {
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
	s.context.CreateFont(name, path)
}

func (s *Surface) MoveTo(x, y float64) {
	s.context.MoveTo(float32(x), float32(y))
}

func (s *Surface) SetFillColor(color uint) {
	r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	s.context.SetFillColor(nanovgo.Color{r, g, b, a})
}

func (s *Surface) SetStrokeColor(color uint) {
	r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	s.context.SetStrokeColor(nanovgo.Color{r, g, b, a})
}

func (s *Surface) SetStrokeWidth(width float64) {
	s.context.SetStrokeWidth(float32(width))
}

func (s *Surface) Stroke() {
	s.context.Stroke()
}

func (s *Surface) Arc(xc, yc, radius, angle1, angle2 float64) {
	// TODO(lbayes): Update external Surface to include direction and facilitate for Cairo
	s.context.Arc(float32(xc), float32(yc), float32(radius), float32(angle1), float32(angle2), nanovgo.Clockwise)
}

func (s *Surface) BeginPath() {
	s.context.BeginPath()
}

func (s *Surface) DebugDumpPathCache() {
	s.context.DebugDumpPathCache()
}

func (s *Surface) Fill() {
	s.context.Fill()
}

func (s *Surface) Rect(x, y, width, height float64) {
	s.context.Rect(float32(x), float32(y), float32(width), float32(height))
}

func (s *Surface) RoundedRect(x, y, width, height, radius float64) {
	s.context.RoundedRect(float32(x), float32(y), float32(width), float32(height), float32(radius))
}

func (s *Surface) SetFontSize(size float64) {
	s.context.SetFontSize(float32(size))
}

func (s *Surface) SetFontFace(face string) {
	s.context.SetFontFace(face)
}

func (s *Surface) Text(x, y float64, text string) {
	// TODO(lbayes): Add validation that ensures required calls have been made before calling this function (e.g., SetFontFace)
	s.context.Text(float32(x), float32(y), text)
}

func (s *Surface) TextBounds(face string, size float64, text string) (x, y, w, h float64) {
	f := s.Font(face)
	f.SetSize(size)

	_, _, h = f.VerticalMetrics()
	mW, bounds := f.Bounds(text)
	return bounds[0], bounds[1], mW, h
}

func (s *Surface) Flags() nanovgo.CreateFlags {
	var result int
	for _, flag := range s.flags {
		result = result | int(flag)
	}
	return nanovgo.CreateFlags(result)
}

func NewSurface(options ...Option) *Surface {
	s := &Surface{}

	for _, option := range options {
		option(s)
	}
	return s
}

func NewWithRoboto(options ...Option) *Surface {
	s := NewSurface(options...)
	s.AddFont("Roboto", "../../third_party/fonts/Roboto/Roboto-Regular.ttf")
	// s.AddFont("Roboto-Light", "../../third_party/fonts/Roboto/Roboto-Light.ttf")
	// s.AddFont("Roboto-Bold", "../../third_party/fonts/Roboto/Roboto-Bold.ttf")
	// s.AddFont("Roboto-Black", "../../third_party/fonts/Roboto/Roboto-Black.ttf")
	return s
}
