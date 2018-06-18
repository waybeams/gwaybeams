package ctrl_test

import (
	"path/filepath"
	"testing"

	"github.com/waybeams/waybeams/pkg/env/fake"
	"github.com/waybeams/waybeams/pkg/layout"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
)

func TestButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		b := ctrl.Button()
		assert.NotNil(b)
		assert.True(b.IsMeasured())
		assert.True(b.IsFocusable())
	})

	t.Run("states", func(t *testing.T) {
		b := ctrl.Button(opts.Text("abcd"))

		assert.Equal(b.State(), "active")
		b.Emit(events.New(events.Entered, b, nil))
		assert.Equal(b.State(), "hovered")
		b.Emit(events.New(events.Pressed, b, nil))
		assert.Equal(b.State(), "pressed")
		b.Emit(events.New(events.Released, b, nil))
		assert.Equal(b.State(), "hovered")
		b.Emit(events.New(events.Exited, b, nil))
		assert.Equal(b.State(), "active")
	})

	t.Run("Label size", func(t *testing.T) {
		b := ctrl.Button(opts.Text("Hello World"))
		layout.Layout(b, fakes.NewSurfaceFrom(filepath.Join("..", "..")))

		// NOTE(lbayes): This test verifies that the Label size is measured
		// properly based on text content and requires the fake surface to
		// be configured with a valid reference to Roboto.ttf.
		assert.Equal(b.Width(), 121)
		assert.Equal(b.Height(), 34)
	})
}
