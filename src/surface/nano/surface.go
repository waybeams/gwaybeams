package nano

import (
	"font"
	"github.com/shibukawa/nanovgo"
	"helpers"
)

type Surface struct {
	context *nanovgo.Context
	flags   []nanovgo.CreateFlags
	width   float64
	height  float64
	fonts   map[string]*font.Font
}

func (s *Surface) Init() {
	context, err := nanovgo.NewContext(s.Flags())
	if err != nil {
		panic(err)
	}

	s.context = context
}

func (s *Surface) Close() {
	if s.context != nil {
		s.context.Delete()
	}
}

func (s *Surface) BeginFrame(w, h float64) {
	s.CreateFonts()

	// ratio := float32(w / h)
	s.context.BeginFrame(int(w), int(h), 1)
}

func (s *Surface) EndFrame() {
	s.context.EndFrame()
}

func (s *Surface) getFonts() map[string]*font.Font {
	if s.fonts == nil {
		s.fonts = make(map[string]*font.Font)
	}
	return s.fonts
}

func (s *Surface) AddFont(name string, path string) {
	fonts := s.getFonts()
	if fonts[name] == nil {
		fonts[name] = font.New(name, path)
	}
}

func (s *Surface) CreateFonts() {
	for _, font := range s.getFonts() {
		if !font.Created {
			s.CreateFont(font.Name, font.Path)
			font.Created = true
		}
	}
}

func (s *Surface) Font(name string) *font.Font {
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

func (s *Surface) MoveTo(x float64, y float64) {
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

func (s *Surface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
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

func (s *Surface) Text(x float64, y float64, text string) {
	// TODO(lbayes): Add validation that ensures required calls have been made before calling this function (e.g., SetFontFace)
	s.context.Text(float32(x), float32(y), text)
}

func (s *Surface) Flags() nanovgo.CreateFlags {
	var result int
	for _, flag := range s.flags {
		result = result | int(flag)
	}
	return nanovgo.CreateFlags(result)
}

type Option func(s *Surface)

func Width(width float64) Option {
	return func(s *Surface) {
		s.SetWidth(width)
	}
}

func Height(height float64) Option {
	return func(s *Surface) {
		s.SetHeight(height)
	}
}

func AntiAlias() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.AntiAlias)
	}
}

func StencilStrokes() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.StencilStrokes)
	}
}

func Debug() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.Debug)
	}
}

func Font(name, path string) Option {
	return func(s *Surface) {
		s.AddFont(name, path)
	}
}

func New(options ...Option) *Surface {
	s := &Surface{}

	for _, option := range options {
		option(s)
	}
	return s
}
