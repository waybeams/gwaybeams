package control

import (
	"assert"
	"testing"
	. "ui"
	"uiold/context"
	"ui/control"
	. "ui/controls"
	. "uiold/opts"
)

func TestCoordToControl(t *testing.T) {
	var createTree = func() Displayable {
		root := VBox(context.New(), ID("root"), Padding(10), Width(100), Height(100), Children(func(c Context) {
			Button(c, ID("abcd"), FlexWidth(1), FlexHeight(1))
			Button(c, ID("efgh"), FlexWidth(1), FlexHeight(1), Padding(5), Children(func() {
				Box(c, ID("efgh.child"), FlexWidth(1), FlexHeight(1))
			}))
			Button(c, ID("ijkl"), FlexWidth(1), FlexHeight(1))
			Button(c, ID("mnop"), FlexWidth(1), FlexHeight(1))
		}))
		return root
	}

	t.Run("CoordToControl", func(t *testing.T) {
		t.Run("Callable", func(t *testing.T) {
			root := Box(context.New(), Width(100), Height(100))
			result := control.CoordToControl(root, 50, 50)
			assert.Equal(t, root, result)
		})

		t.Run("ContainsCoordinate", func(t *testing.T) {
			instance := Button(context.New(), Width(100), Height(100))
			assert.True(t, control.ContainsCoordinate(instance, 20, 20))
		})

		t.Run("returns root when out of bounds lower right", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 1000, 1000)
			assert.Equal(t, root.ID(), result.ID())
		})

		t.Run("returns root when out of bounds upper left", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, -1000, -1000)
			assert.Equal(t, root.ID(), result.ID())
		})

		t.Run("Returns element within bounds", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 15, 15)
			assert.NotNil(t, result)
			assert.Equal(t, result.Path(), "/root/abcd")
		})

		t.Run("Returns element on first pixel", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 10, 10)
			assert.NotNil(t, result)
			assert.Equal(t, result.ID(), "abcd")
		})

		t.Run("Returns element on last pixel", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 80, 20)
			assert.NotNil(t, result)
			assert.Equal(t, result.ID(), "abcd")
		})

		t.Run("Returns next element", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 20, 35)
			assert.NotNil(t, result)
			assert.Equal(t, result.ID(), "efgh")
		})

		t.Run("Only returns Focusable elements", func(t *testing.T) {
			root := createTree()
			result := control.CoordToControl(root, 20, 40)
			assert.NotNil(t, result)
			assert.Equal(t, result.ID(), "efgh", "NOT efgh.child")
		})
	})

	t.Run("LocalToGlobal", func(t *testing.T) {
		t.Run("No parent", func(t *testing.T) {
			instance := Button(context.New(), Width(100), Height(100))
			x, y := control.LocalToGlobal(instance, 50, 60)
			assert.Equal(t, x, 50)
			assert.Equal(t, y, 60)
		})

		t.Run("Single parent", func(t *testing.T) {
			instance := VBox(context.New(), Padding(10), Width(100), Height(100), Children(func(c Context) {
				Button(c, ID("abcd"), Width(100), Height(50))
			}))
			abcd := instance.FindControlById("abcd")
			x, y := control.LocalToGlobal(abcd, 20, 30)
			assert.Equal(t, x, 30)
			assert.Equal(t, y, 40)
		})

		t.Run("Nested Parents", func(t *testing.T) {
			root := VBox(context.New(), ID("root"), Padding(10), Width(100), Height(100), Children(func(c Context) {
				VBox(c, ID("abcd"), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
					VBox(c, ID("efgh"), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
						VBox(c, ID("ijkl"), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
							Button(c, ID("mnop"), Width(40), Height(40))
						}))
					}))
				}))
			}))

			x, y := control.LocalToGlobal(root, 20, 30)
			assert.Equal(t, x, 20)
			assert.Equal(t, y, 30)

			abcd := root.FindControlById("abcd")
			x, y = control.LocalToGlobal(abcd, 20, 30)
			assert.Equal(t, x, 30)
			assert.Equal(t, y, 40)

			efgh := root.FindControlById("efgh")
			x, y = control.LocalToGlobal(efgh, 20, 30)
			assert.Equal(t, x, 40)
			assert.Equal(t, y, 50)

			ijkl := root.FindControlById("ijkl")
			x, y = control.LocalToGlobal(ijkl, 20, 30)
			assert.Equal(t, x, 50)
			assert.Equal(t, y, 60)

			mnop := root.FindControlById("mnop")
			x, y = control.LocalToGlobal(mnop, 20, 30)
			assert.Equal(t, x, 60)
			assert.Equal(t, y, 70)
		})
	})

}
