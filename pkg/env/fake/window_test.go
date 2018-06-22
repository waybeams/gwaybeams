package fake_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/fake"
)

func TestWinFake(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		w := fake.NewFakeWindow()
		w.SetWidth(30)
		w.SetHeight(40)
		assert.Equal(w.Width(), 30)
		assert.Equal(w.Height(), 40)
	})
}
