package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/examples/todo/ctrl"
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestItemSpec(t *testing.T) {

	t.Run("Displays model state", func(t *testing.T) {
		m := model.New()
		m.CreateItem("Item One")

		itemModel := m.CurrentItems()[0]
		s := ctrl.ItemSpec(itemModel, 2)

		desc := spec.FirstByKey(s, "desc")
		assert.Equal(desc.Text(), "Item One")

		toggle := spec.FirstByKey(s, "btn")
		assert.Equal(toggle.Text(), "[  ]")

		// Click the toggle completed button
		toggle.Emit(events.New(events.Clicked, toggle, nil))

		// Manually build a new component from the updated model state
		s = ctrl.ItemSpec(itemModel, 2)
		toggle = spec.FirstByKey(s, "btn")
		assert.Equal(toggle.Text(), "[X]")
	})
}
