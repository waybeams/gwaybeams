package browser_test

import (
	"testing"

	jsCanvas "github.com/oskca/gopherjs-canvas"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/browser"
	"github.com/waybeams/waybeams/pkg/spec"
)

type FakeExternalCanvas struct {
}

func (f *FakeExternalCanvas) Set(key string, value interface{}) {
}

func (f *FakeExternalCanvas) GetContext2D() *jsCanvas.Context2D {
	return nil
}

func createSurface() spec.Surface {
	c := &FakeExternalCanvas{}
	s := browser.NewSurface(c)
	return s
}

func TestBrowserSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		s := createSurface()
		assert.NotNil(s)
	})
}
