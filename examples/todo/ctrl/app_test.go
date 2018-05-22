package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"testing"
)

func TestAppControl(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		m := model.New()
		render := ctrl.AppRenderer(m)
		assert.Equal(render().Children()[0].ChildCount(), 4)
	})

	t.Run("Children created", func(t *testing.T) {
		m := model.New()
		m.CreateItem("Item One")
		m.CreateItem("Item Two")
		m.CreateItem("Item Three")
		m.CreateItem("Item Four")
		m.CreateItem("Item Five")
		tree := ctrl.AppRenderer(m)()
		items := tree.Children()[0].Children()[2].Children()
		assert.Equal(len(items), 5)
		assert.Equal(items[0].Children()[1].Text(), "Item One")
	})
}
