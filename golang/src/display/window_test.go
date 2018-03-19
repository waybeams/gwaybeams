package display

import (
	"assert"
	"testing"
)

func TestWindow(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		win := NewWindow(&Opts{})
		assert.NotNil(win)
	})
}
