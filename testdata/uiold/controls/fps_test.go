package controls

import (
	"github.com/waybeams/assert"
	"uiold/context"
	"testing"
)

func TestFps(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := FPS(context.New())
		assert.NotNil(instance)
	})
}
