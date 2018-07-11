package builder_test

import (
	"testing"
	"time"

	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/spec"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/env/fake"
)

func TestBuilder(t *testing.T) {

	t.Run("Surface", func(t *testing.T) {
		factoryCalled := false
		fakeWindow := fake.NewWindow()
		fakeSurface := fake.NewSurface()
		fakeAppFactory := func() spec.ReadWriter {
			factoryCalled = true
			return ctrl.VBox()
		}

		b := builder.New(fakeWindow, fakeSurface, fakeAppFactory)

		// Ensure we close the blocked goroutine.
		defer b.Close()
		// Listen in a goroutine.
		go b.Listen()
		// Move time forward and ensure our factory was called.
		fakeWindow.Clock().Add(100 * time.Millisecond)
		assert.True(factoryCalled)
	})
}
