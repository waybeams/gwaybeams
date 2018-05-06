package helpers

import (
	"testing"
)

func TestColors(t *testing.T) {

	t.Run("HexIntToRgba", func(t *testing.T) {
		t.Run("Orange", func(t *testing.T) {
			r, g, b, a := HexIntToRgba(0xffcc00ff)

			if r != 0xff {
				t.Errorf("Expected Red (%d) to equal 0xff", r)
			}

			if g != 0xcc {
				t.Errorf("Expected Green (%d) to equal 0xcc", g)
			}

			if b != 0x00 {
				t.Errorf("Expected Blue (%d) to equal 0x00", b)
			}

			if a != 0xff {
				t.Errorf("Expected Alpha (%d) to equal 0xff", a)
			}
		})

		t.Run("White", func(t *testing.T) {
			r, g, b, a := HexIntToRgba(0xffffffff)

			if r != 0xff {
				t.Errorf("Expected Red (%d) to equal 0xff", r)
			}

			if g != 0xff {
				t.Errorf("Expected Green (%d) to equal 0xff", g)
			}

			if b != 0xff {
				t.Errorf("Expected Blue (%d) to equal 0xff", b)
			}

			if a != 0xff {
				t.Errorf("Expected Alpha (%d) to equal 0xff", a)
			}
		})
	})

	t.Run("HexIntToRgbaFloat64", func(t *testing.T) {
		r, g, b, a := HexIntToRgbaFloat64(0xffcc00ff)

		if r != 1.0 {
			t.Errorf("Expected Red (%f) to equal 1.0", r)
		}

		if g != 0.8 {
			t.Errorf("Expected Green (%f) to equal 0.8", g)
		}

		if b != 0.0 {
			t.Errorf("Expected Blue (%f) to equal 0.0", b)
		}

		if a != 1.0 {
			t.Errorf("Expected Alpha (%f) to equal 1.0", a)
		}
	})

	t.Run("HexIntToRgb", func(t *testing.T) {
		t.Run("black", func(t *testing.T) {
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

		t.Run("blue", func(t *testing.T) {
			r, g, b := HexIntToRgb(0x0000ff)

			if r != 0x00 {
				t.Errorf("Expected Red (%d) to equal 0x00", r)
			}

			if g != 0x00 {
				t.Errorf("Expected Green (%d) to equal 0x00", g)
			}

			if b != 0xff {
				t.Errorf("Expected Blue (%d) to equal 0xff", b)
			}
		})
	})

	t.Run("HexIntToRgbFloat64", func(t *testing.T) {
		t.Run("Orange", func(t *testing.T) {
			r, g, b := HexIntToRgbFloat64(0xffcc00)

			if r != 1.0 {
				t.Errorf("Expected Red (%f) to equal 1.0", r)
			}

			if g != 0.8 {
				t.Errorf("Expected Green (%f) to equal 0.8", g)
			}

			if b != 0.0 {
				t.Errorf("Expected Blue (%f) to equal 0.0", b)
			}
		})
	})
}
