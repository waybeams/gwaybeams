package model_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/examples/todo/model"
	"testing"
)

func TestItemModel(t *testing.T) {
	t.Run("Description", func(t *testing.T) {
		m := model.New()
		item := m.CreateItem("Item One")
		assert.Equal(item.Description, "Item One")
	})

	t.Run("ToggleCompleted", func(t *testing.T) {
		m := model.New()
		item := m.CreateItem("Item One")
		assert.False(item.IsCompleted(), "Should NOT be completed 1")

		item.ToggleCompleted()
		assert.True(item.IsCompleted(), "Should be completed 2")

		item.ToggleCompleted()
		assert.False(item.IsCompleted(), "Should NOT be completed 3")
	})

	t.Run("Delete item", func(t *testing.T) {
		m := model.New()
		assert.Equal(len(m.CurrentItems()), 0)

		item := m.CreateItem("Item One")
		assert.Equal(len(m.CurrentItems()), 1)

		item.Delete()
		assert.Equal(len(m.CurrentItems()), 0)
	})
}
