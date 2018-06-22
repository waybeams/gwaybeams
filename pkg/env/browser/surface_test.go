package browser_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/browser"
	"github.com/waybeams/waybeams/pkg/spec"
)

func createSurface() spec.Surface {
	// canvas := fake.NewBrowserCanvas()
	s := browser.NewSurface(nil)
	return s
}

func TestBrowserSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		s := createSurface()
		assert.NotNil(s)
	})
}
