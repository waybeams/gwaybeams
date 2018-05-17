package todomvc

import (
	"github.com/waybeams/assert"
	"clock"
	"testing"
	"todomvc"
)

func createTodoAppModel() *todomvc.TodoAppModel {
	m := &todomvc.TodoAppModel{Clock: clock.NewFake()}
	m.PushItem("abcd")
	m.PushItem("efgh")
	m.PushItem("ijkl")
	m.PushItem("mnop")

	return m
}

func TestTodoMVCModels(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		m := &todomvc.TodoAppModel{}
		assert.NotNil(t, m)
	})

	t.Run("PushItem", func(t *testing.T) {
		m := createTodoAppModel()
		items := m.AllItems()
		assert.Equal(t, len(items), 4)
		assert.Equal(t, items[0].Text, "abcd")
		assert.Equal(t, items[1].Text, "efgh")
		assert.Equal(t, items[2].Text, "ijkl")
		assert.Equal(t, items[3].Text, "mnop")
	})

	t.Run("CompleteItem", func(t *testing.T) {
		m := createTodoAppModel()
		m.ItemAt(1).Complete()
		completed := m.CompletedItems()
		assert.Equal(t, len(completed), 1)
		assert.Equal(t, completed[0].Text, "efgh")

		pending := m.PendingItems()
		assert.Equal(t, len(pending), 3)
		assert.Equal(t, pending[0].Text, "abcd")
		assert.Equal(t, pending[1].Text, "ijkl")
		assert.Equal(t, pending[2].Text, "mnop")
	})

	t.Run("Trims input strings", func(t *testing.T) {
		m := createTodoAppModel()
		m.PushItem(" wxyz ")
		item := m.LastItem()
		assert.Equal(t, item.Text, "wxyz")
	})

	t.Run("Update item", func(t *testing.T) {
		m := createTodoAppModel()
		m.UpdateItemAt(0, " wxyz ")
		item := m.ItemAt(0)
		assert.Equal(t, item.Text, "wxyz")
	})

	t.Run("Pending Label", func(t *testing.T) {
		m := createTodoAppModel()
		m.ItemAt(0).Complete()

		// 2 or more items remaining
		label := m.PendingLabel()
		assert.Equal(t, label, "3 items left")

		// 1 item remaining
		m.ItemAt(1).Complete()
		m.ItemAt(2).Complete()
		label = m.PendingLabel()
		assert.Equal(t, label, "1 item left")

		// 0 items remaining
		m.ItemAt(3).Complete()
		label = m.PendingLabel()
		assert.Equal(t, label, "0 items left")
	})

	t.Run("FilterSelection", func(t *testing.T) {
		m := createTodoAppModel()
		assert.Equal(t, m.FilterSelection(), "All")
	})

	t.Run("Update Filterselection", func(t *testing.T) {
		m := createTodoAppModel()
		m.SetFilterSelection("Active")
		assert.Equal(t, m.FilterSelection(), "Active")
	})
}
