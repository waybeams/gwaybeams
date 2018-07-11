package fake_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/fake"
)

func TestFakeSurface(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		s := fake.NewSurface()
		assert.NotNil(s)
	})

	t.Run("Width", func(t *testing.T) {
		s := fake.NewSurface()
		s.SetWidth(23)
		assert.Equal(s.Width(), 23)
		cmds := s.GetCommands()
		assert.Equal(len(cmds), 2)
		assert.Equal(cmds[0].Name, "SetWidth")
		assert.Equal(cmds[0].Args[0], 23)
		assert.Equal(cmds[1].Name, "Width")
		assert.Equal(len(cmds[1].Args), 0)
	})

	t.Run("Height", func(t *testing.T) {
		s := fake.NewSurface()
		s.SetHeight(24)
		assert.Equal(s.Height(), 24)
		cmds := s.GetCommands()
		assert.Equal(len(cmds), 2)
		assert.Equal(cmds[0].Name, "SetHeight")
		assert.Equal(cmds[0].Args[0], 24)
		assert.Equal(cmds[1].Name, "Height")
		assert.Equal(len(cmds[1].Args), 0)
	})
}
