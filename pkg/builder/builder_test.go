package builder_test

import (
	"testing"
	"time"

	"github.com/waybeams/waybeams/pkg/clock"

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
		fakeClock := clock.NewFake()

		b := builder.New(fakeWindow, fakeSurface, fakeAppFactory, fakeClock)

		// Ensure we close the blocked goroutine.
		defer b.Close()
		// Listen in a goroutine.
		go b.Listen()
		// Move time forward and ensure our factory was called.
		fakeClock.Add(100 * time.Millisecond)
		assert.True(factoryCalled)
	})
}
