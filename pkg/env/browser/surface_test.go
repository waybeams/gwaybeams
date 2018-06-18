package browser_test

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/browser"
)

func TestBrowserSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		elem := &js.Object{}
		w := browser.NewSurface(browser.Canvas(elem))
		assert.NotNil(w)
	})
}
