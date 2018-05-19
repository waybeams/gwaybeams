package controls_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/controls"
	"github.com/waybeams/waybeams/examples/todo/model"
	"testing"
)

func TestAppControl(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		m := model.New()
		render := controls.AppRenderer(m)
		assert.Equal(render().Children()[0].ChildCount(), 4)
	})
}
