package clock_test

import (
	"testing"
	"time"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/clock"
)

func TestFrameRate(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		fakeClock := clock.NewFake()

		callCount := 0
		var handler = func() bool {
			callCount++
			return false
		}

		// launch the blocking OnFrame call in a go routine so that we can
		// more easily make assertions about it's execution. This is NOT
		// how it should be used.
		go clock.OnFrame(handler, 2, fakeClock)

		assert.Equal(callCount, 0, "Should not be called right away")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 1, "callCount 1")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 2, "callCount 2")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 3, "callCount 3")
	})
}
