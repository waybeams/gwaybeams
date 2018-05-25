package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestFooterSpec(t *testing.T) {
	var createModel = func() *model.App {
		m := model.New()
		m.CreateItem("Item One")
		m.CreateItem("Item Two")
		m.CreateItem("Item Three")
		m.CreateItem("Item Four")
		m.CreateItem("Item Five")

		items := m.ActiveItems()
		// Mark two of the items as completed.
		items[2].ToggleCompleted()
		items[3].ToggleCompleted()
		return m
	}

	t.Run("Default Configuration", func(t *testing.T) {
		footer := ctrl.Footer(createModel(), ctrl.CreateStyles())
		label := spec.FirstByKey(footer, "Item Count")
		assert.Equal(label.Text(), "5 items")

		btn := spec.FirstByKey(footer, "All Button")
		assert.Equal(btn.StrokeSize(), 1)

		btn = spec.FirstByKey(footer, "Active Button")
		assert.Equal(btn.StrokeSize(), 0)

		btn = spec.FirstByKey(footer, "Completed Button")
		assert.Equal(btn.StrokeSize(), 0)

		btn = spec.FirstByKey(footer, "Clear Completed Button")
		assert.Equal(btn.StrokeSize(), 0)
	})

	t.Run("Clicks", func(t *testing.T) {
		m := createModel()
		footer := ctrl.Footer(m, ctrl.CreateStyles())

		assert.Equal(m.Showing(), model.AllItems, "Model shows All Items by default")

		btn := spec.FirstByKey(footer, "Active Button")
		btn.Bubble(events.New(events.Clicked, btn, nil))

		assert.Equal(m.Showing(), model.ActiveItems, "Active Items button changes model selection")
	})
}
