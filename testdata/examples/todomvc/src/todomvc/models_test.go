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
		assert.NotNil(m)
	})

	t.Run("PushItem", func(t *testing.T) {
		m := createTodoAppModel()
		items := m.AllItems()
		assert.Equal(len(items), 4)
		assert.Equal(items[0].Text, "abcd")
		assert.Equal(items[1].Text, "efgh")
		assert.Equal(items[2].Text, "ijkl")
		assert.Equal(items[3].Text, "mnop")
	})

	t.Run("CompleteItem", func(t *testing.T) {
		m := createTodoAppModel()
		m.ItemAt(1).Complete()
		completed := m.CompletedItems()
		assert.Equal(len(completed), 1)
		assert.Equal(completed[0].Text, "efgh")

		pending := m.PendingItems()
		assert.Equal(len(pending), 3)
		assert.Equal(pending[0].Text, "abcd")
		assert.Equal(pending[1].Text, "ijkl")
		assert.Equal(pending[2].Text, "mnop")
	})

	t.Run("Trims input strings", func(t *testing.T) {
		m := createTodoAppModel()
		m.PushItem(" wxyz ")
		item := m.LastItem()
		assert.Equal(item.Text, "wxyz")
	})

	t.Run("Update item", func(t *testing.T) {
		m := createTodoAppModel()
		m.UpdateItemAt(0, " wxyz ")
		item := m.ItemAt(0)
		assert.Equal(item.Text, "wxyz")
	})

	t.Run("Pending Label", func(t *testing.T) {
		m := createTodoAppModel()
		m.ItemAt(0).Complete()

		// 2 or more items remaining
		label := m.PendingLabel()
		assert.Equal(label, "3 items left")

		// 1 item remaining
		m.ItemAt(1).Complete()
		m.ItemAt(2).Complete()
		label = m.PendingLabel()
		assert.Equal(label, "1 item left")

		// 0 items remaining
		m.ItemAt(3).Complete()
		label = m.PendingLabel()
		assert.Equal(label, "0 items left")
	})

	t.Run("FilterSelection", func(t *testing.T) {
		m := createTodoAppModel()
		assert.Equal(m.FilterSelection(), "All")
	})

	t.Run("Update Filterselection", func(t *testing.T) {
		m := createTodoAppModel()
		m.SetFilterSelection("Active")
		assert.Equal(m.FilterSelection(), "Active")
	})
}
