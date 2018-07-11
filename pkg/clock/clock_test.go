package clock_test

import (
	"testing"
	"time"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/clock"
)

const NinetyFive = 800000000000000000

func TestClock(t *testing.T) {

	t.Run("Real Clock instantiable", func(t *testing.T) {
		c := clock.New()
		n := c.Now()
		assert.NotNil(n)
	})

	t.Run("Set", func(t *testing.T) {
		c := clock.NewFake()
		c.Set(time.Unix(0, NinetyFive))
		n := c.Now()
		assert.Match("1995-05", n.String())
	})

	t.Run("Add", func(t *testing.T) {
		c := clock.NewFake()
		c.Set(time.Unix(0, NinetyFive))

		c.Add(8760 * time.Hour) // One year in hours
		n := c.Now()
		assert.Match("1996-05", n.String())
	})

	t.Run("OnFrame", func(t *testing.T) {
		fakeClock := clock.NewFake()

		callCount := 0
		var handler = func() bool {
			callCount++
			return false
		}

		// launch the blocking OnFrame call in a go routine so that we can
		// more easily make assertions about it's execution. This is NOT
		// how it should be used.
		go fakeClock.OnFrame(handler, 2)

		assert.Equal(callCount, 0, "Should not be called right away")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 1, "callCount 1")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 2, "callCount 2")
		fakeClock.Add(500 * time.Millisecond)
		assert.Equal(callCount, 3, "callCount 3")
	})
}
