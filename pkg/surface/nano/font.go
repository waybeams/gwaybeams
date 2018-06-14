package nano

import (
	"github.com/waybeams/waybeams/pkg/spec"

	fsm "github.com/shibukawa/nanovgo/fontstashmini"
)

const nvgInitFontImageSize = 512

type Font struct {
	name    string
	path    string
	created bool
	stash   *fsm.FontStash
}

func (f *Font) Name() string {
	return f.name
}

func (f *Font) Path() string {
	return f.path
}

func (f *Font) OnCreated() {
	f.created = true
}

func (f *Font) IsCreated() bool {
	return f.created
}

func (f *Font) getStash() *fsm.FontStash {
	if f.stash == nil {
		f.stash = fsm.New(nvgInitFontImageSize, nvgInitFontImageSize)
		result := f.stash.AddFont(f.name, f.path)
		if result == -1 {
			msg := "Unable to add font, likely bad path: " + f.path
			panic(msg)
		}
	}
	return f.stash
}

func (f *Font) SetSize(size float32) {
	f.getStash().SetSize(size)
}

func (f *Font) SetAlign(align spec.Alignment) {
	var fsa fsm.FONSAlign
	switch align {
	case spec.AlignBottom:
		fsa = fsm.ALIGN_BOTTOM
	case spec.AlignCenter:
		fsa = fsm.ALIGN_CENTER
	case spec.AlignLeft:
		fsa = fsm.ALIGN_LEFT
	case spec.AlignRight:
		fsa = fsm.ALIGN_RIGHT
	case spec.AlignTop:
		fsa = fsm.ALIGN_TOP
	case spec.AlignMiddle:
		fsa = fsm.ALIGN_MIDDLE

	}
	f.getStash().SetAlign(fsa)
}

func (f *Font) VerticalMetrics() (ascender, descender, lineHeight float32) {
	stash := f.getStash()
	return stash.VerticalMetrics()
}

func (f *Font) Bounds(value string) (width float32, bounds []float32) {
	stash := f.getStash()
	return stash.TextBounds(0, 0, value)
}

func NewFont(name string, path string) *Font {
	return &Font{
		name: name,
		path: path,
	}
}
