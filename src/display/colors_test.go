package display

import (
	"testing"
)

func TestRgbaToParts(t *testing.T) {
	t.Run("Handles black", func(t *testing.T) {
		r, g, b := HexIntToRgb(0xffcc11)

		if r != 0xff {
			t.Errorf("Expected Red (%d) to equal 0xff", r)
		}

		if g != 0xcc {
			t.Errorf("Expected Green (%d) to equal 0xcc", g)
		}

		if b != 0x11 {
			t.Errorf("Expected Blue (%d) to equal 0x00", b)
		}
	})
}
