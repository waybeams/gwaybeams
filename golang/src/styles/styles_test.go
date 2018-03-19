package styles

import (
	"assert"
	"testing"
)

func TestStyles(t *testing.T) {
	t.Run("Empty Styles", func(t *testing.T) {
		assert.NotNil(Styles())
	})

	t.Run("Single Style", func(t *testing.T) {
		styles := Styles(BgColor(0xfc0))
		assert.Equal(styles["BgColor"].(uint), uint(0xfc0))
	})

	t.Run("Multiple Styles", func(t *testing.T) {
		styles := Styles(BgColor(0xfc0), BorderColor(0x0cf), BorderSize(2))
		assert.Equal(styles["BgColor"].(uint), uint(0xfc0))
		assert.Equal(styles["BorderColor"].(uint), uint(0x0cf))
		assert.Equal(styles["BorderSize"].(int), int(2))
	})
}
