package controls_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/controls"
	"github.com/waybeams/waybeams/examples/todo/model"
	"testing"
)

func TestItemSpec(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		m := &model.Item{Description: "Item One"}
		s := controls.ItemSpec(m, 2)
		assert.NotNil(s)
	})
}
