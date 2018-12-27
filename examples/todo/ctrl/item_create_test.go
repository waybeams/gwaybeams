package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/examples/todo/ctrl"
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestItemCreate(t *testing.T) {

	var createModel = func() *model.App {
		m := model.NewSample()
		items := m.CurrentItems()
		items[2].ToggleCompleted()
		items[3].ToggleCompleted()
		return m
	}

	t.Run("Instantiable", func(t *testing.T) {
		model := createModel()
		instance := ctrl.ItemCreate(model)
		assert.NotNil(instance)
	})

	t.Run("Allows option overrides", func(t *testing.T) {
		model := createModel()
		instance := ctrl.ItemCreate(model, opts.FlexWidth(3))
		assert.Equal(instance.FlexWidth(), 3)
	})

	t.Run("Submits on enter", func(t *testing.T) {
		model := createModel()
		instance := ctrl.ItemCreate(model)

		allItems := model.AllItems()

		assert.Equal(len(allItems), 6)

		textInput := spec.FirstByKey(instance, ctrl.NewItemInput)
		textInput.SetText("Hello World")

		instance.Bubble(events.New(events.EnterKeyReleased, instance, nil))

		assert.Equal(len(model.AllItems()), 7)
		allItems = model.AllItems()

		lastItem := allItems[len(allItems)-1]
		assert.Equal(lastItem.Description, "Hello World")
	})
}
