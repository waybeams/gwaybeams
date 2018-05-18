package fakes_test

import (
	"github.com/waybeams/assert"
	"fakes"
	"opts"
	"testing"
)

func TestFakes(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		ctrl := fakes.Fake(opts.Width(20), opts.Height(34))
		assert.Equal(ctrl.Width(), 20)
		assert.Equal(ctrl.Height(), 34)
	})

	t.Run("Container", func(t *testing.T) {
		ctr := fakes.FakeContainer()
		assert.Equal(ctr.ChildCount(), 3)
	})
}
