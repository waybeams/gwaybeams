package browser_test

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
	jsCanvas "github.com/oskca/gopherjs-canvas"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/browser"
	"github.com/waybeams/waybeams/pkg/spec"
)

func createSurface() spec.Surface {
	c := jsCanvas.New(&js.Object{})
	s := browser.NewSurface(c)
	// s.Init()
	return s
}

func TestBrowserSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		s := createSurface()
		assert.NotNil(s)
	})
}
