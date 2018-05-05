package context_test

import (
	"../context" // um, why?!
	"assert"
	"clock"
	"display"
	"testing"
)

func TestContext(t *testing.T) {

	t.Run("Instantiated", func(t *testing.T) {
		c := context.New()
		assert.NotNil(t, c)
	})

	t.Run("Sets defaults", func(t *testing.T) {
		c := context.New()
		assert.NotNil(t, c.Clock())
		assert.NotNil(t, c.Builder())
	})

	t.Run("Accepts Builder", func(t *testing.T) {
		b := display.NewBuilder()
		c := context.New(context.Builder(b))
		assert.Equal(t, b, c.Builder())
	})

	t.Run("Accepts Clock", func(t *testing.T) {
		ck := clock.NewFake()
		c := context.New(context.Clock(ck))
		assert.Equal(t, ck, c.Clock())
	})

	t.Run("Accepts Clock and Builder", func(t *testing.T) {
		ck := clock.NewFake()
		b := display.NewBuilder()
		c := context.New(context.Builder(b), context.Clock(ck))
		assert.Equal(t, c.Clock(), ck)
		assert.Equal(t, c.Builder(), b)
	})
}
