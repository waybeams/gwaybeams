package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestItemList(t *testing.T) {
	var createModel = func() *model.App {
		m := model.NewSample()
		items := m.CurrentItems()
		items[2].ToggleCompleted()
		items[3].ToggleCompleted()
		return m
	}

	t.Run("Configures Children", func(t *testing.T) {
		model := createModel()
		instance := ctrl.ItemList(model)
		kids := instance.Children()
		child := spec.FirstByKey(kids[0], "btn")
		assert.Equal(child.Text(), "[  ]", "Expected btn label NOT to be selected")

		child = spec.FirstByKey(kids[2], "btn")
		assert.Equal(child.Text(), "[X]", "Expected btn label to be selected")
	})

	t.Run("Accepts option overrides", func(t *testing.T) {
		instance := ctrl.ItemList(createModel(), opts.FlexWidth(5))
		assert.Equal(instance.FlexWidth(), 5)
	})
}
