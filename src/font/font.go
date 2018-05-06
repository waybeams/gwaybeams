package font

import (
	"fmt"
	fsm "github.com/shibukawa/nanovgo/fontstashmini"
)

const nvgInitFontImageSize = 512

type Font struct {
	Name    string
	Path    string
	Created bool
	stash   *fsm.FontStash
}

func (f *Font) getStash() *fsm.FontStash {
	if f.stash == nil {
		f.stash = fsm.New(nvgInitFontImageSize, nvgInitFontImageSize)
		result := f.stash.AddFont(f.Name, f.Path)
		if result == -1 {
			msg := fmt.Sprintf("Unable to add font, likely bad path: %v", f.Path)
			panic(msg)
		}
	}
	return f.stash
}

func (f *Font) SetSize(size float32) {
	f.getStash().SetSize(size)
}

func (f *Font) SetAlign(align fsm.FONSAlign) {
	f.getStash().SetAlign(align)
}

func (f *Font) Bounds(value string) (width float32, bounds []float32) {
	stash := f.getStash()
	return stash.TextBounds(0, 0, value)
}

func New(name string, path string) *Font {
	return &Font{
		Name: name,
		Path: path,
	}
}
