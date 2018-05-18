package font_test

import (
	"github.com/waybeams/assert"
	"font"
	"testing"
)

const RobotoTestPath = "../../testdata/Roboto-Regular.ttf"

func TestFont(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance := font.New("abcd", "foo.ttf")
		assert.NotNil(instance)
	})

	t.Run("Loads font only when requested", func(t *testing.T) {
		instance := font.New("abcd", RobotoTestPath)
		instance.SetSize(18)
		w, bounds := instance.Bounds("abcd")

		assert.Equal(w, 34)
		assert.Equal(len(bounds), 4)
		assert.Equal(bounds[0], -1)
		assert.Equal(bounds[1], -13)
		assert.Equal(bounds[2], 34)
		assert.Equal(bounds[3], 2)

		// Change size and verify values are different.
		instance.SetSize(12)
		w, bounds = instance.Bounds("abcd")

		assert.Equal(w, 23)
		assert.Equal(len(bounds), 4)
		assert.Equal(bounds[0], -1)
		assert.Equal(bounds[1], -9)
		assert.Equal(bounds[2], 24)
		assert.Equal(bounds[3], 2)
	})
}
