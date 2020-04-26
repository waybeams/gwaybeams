package model_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/model"
	"testing"
)

func TestAppModel(t *testing.T) {

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

	t.Run("AllItems", func(t *testing.T) {
		m := createModel()
		assert.Equal(len(m.AllItems()), 5)
		m.ShowAllItems()
		assert.Equal(len(m.CurrentItems()), 5)
	})

	t.Run("ActiveItems", func(t *testing.T) {
		m := createModel()
		assert.Equal(len(m.ActiveItems()), 3)
		m.ShowActiveItems()
		assert.Equal(len(m.CurrentItems()), 3)
	})

	t.Run("CompletedItems", func(t *testing.T) {
		m := createModel()
		assert.Equal(len(m.CompletedItems()), 2)
		m.ShowCompletedItems()
		assert.Equal(len(m.CurrentItems()), 2)
	})

	t.Run("Create item", func(t *testing.T) {
		m := createModel()
		m.CreateItem("Item Six")
		assert.Equal(len(m.AllItems()), 6)
	})

	t.Run("Delete item", func(t *testing.T) {
		m := createModel()
		items := m.CurrentItems()
		m.DeleteItem(items[0])
		assert.Equal(len(m.CurrentItems()), 4)
	})

	t.Run("Complete item", func(t *testing.T) {
		m := createModel()
		assert.Equal(len(m.CompletedItems()), 2, "Two completed items")

		items := m.ActiveItems()
		items[0].ToggleCompleted()
		assert.Equal(len(m.CompletedItems()), 3, "We've completed another item")

		items[0].ToggleCompleted()
		assert.Equal(len(m.CompletedItems()), 2, "We've un-completed that item")
	})

	t.Run("No more completed items flips current items", func(t *testing.T) {
		m := createModel()
		m.ShowCompletedItems()

		items := m.CurrentItems()
		items[0].ToggleCompleted()
		assert.Equal(m.Showing(), model.CompletedItems, "showing is unchanged")

		items[1].ToggleCompleted()
		assert.Equal(m.Showing(), model.AllItems, "Automatically flips state to all items")
	})

	t.Run("No more active items flips current items", func(t *testing.T) {
		m := createModel()
		m.ShowActiveItems()

		items := m.CurrentItems()
		items[0].ToggleCompleted()
		items[1].ToggleCompleted()
		assert.Equal(m.Showing(), model.ActiveItems, "showing is unchanged")

		items[2].ToggleCompleted()
		assert.Equal(m.Showing(), model.AllItems, "Automatically flips state to all items")
	})
}
