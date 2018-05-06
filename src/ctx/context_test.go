package ctx_test

import (
	"assert"
	"clock"
	"ctx"
	"testing"
	"ui"
)

func TestContext(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		c := ctx.New()
		assert.NotNil(t, c)
	})

	t.Run("Sets defaults", func(t *testing.T) {
		c := ctx.New()
		assert.NotNil(t, c.Clock())
		assert.NotNil(t, c.Builder())
	})

	t.Run("Accepts Builder", func(t *testing.T) {
		b := ui.NewBuilder()
		c := ctx.New(ctx.Builder(b))
		assert.Equal(t, b, c.Builder())
	})

	t.Run("Accepts Clock", func(t *testing.T) {
		ck := clock.NewFake()
		c := ctx.New(ctx.Clock(ck))
		assert.Equal(t, ck, c.Clock())
	})

	t.Run("Accepts Clock and Builder", func(t *testing.T) {
		ck := clock.NewFake()
		b := ui.NewBuilder()
		c := ctx.New(ctx.Builder(b), ctx.Clock(ck))
		assert.Equal(t, c.Clock(), ck)
		assert.Equal(t, c.Builder(), b)
	})
}
