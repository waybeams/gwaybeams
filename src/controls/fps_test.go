package controls

import (
	"assert"
	"ctx"
	"testing"
)

func TestFps(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := FPS(ctx.New())
		assert.NotNil(t, instance)
	})
}
