package display

import (
	"assert"
	"testing"
)

func TestLayoutable(t *testing.T) {
	t.Run("Default Size", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		assert.Equal(t, box.FixedWidth(), -1, "FixedWidth")
		assert.Equal(t, box.FixedHeight(), -1, "FixedHeight")
	})

	t.Run("Default Size after Layout", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		if box.Width() != 0 {
			t.Errorf("Expected width to be 0 but was %v", box.Width())
		}
		if box.Height() != 0 {
			t.Errorf("Expected height to be 0 but was %v", box.HAlign())

		}
	})

	t.Run("GetLayoutType default value", func(t *testing.T) {
		root, _ := Box(NewBuilder())
		if root.LayoutType() != StackLayoutType {
			t.Errorf("Expected %v but got %v", StackLayoutType, root.LayoutType())
		}
	})

	t.Run("MaxHeight constrained Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Height(51), MaxHeight(41))
		assert.Equal(t, box.Height(), 41.0)
	})

	t.Run("MaxWidth constrained Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(50), MaxWidth(40))
		assert.Equal(t, box.Width(), 40.0)
	})

	t.Run("MinHeight becomes unset Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinHeight(20))
		assert.Equal(t, box.Height(), 20.0)
	})

	t.Run("MinWidth becomes unset Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(20))
		assert.Equal(t, box.Width(), 20.0)
	})

	t.Run("MinHeight replaces existing Height", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.SetHeight(10)
		box.SetMinHeight(20)
		assert.Equal(t, box.Height(), 20.0)
	})

	t.Run("MinWidth replaces existing Width", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.SetWidth(10)
		box.SetMinWidth(20)
		assert.Equal(t, box.Width(), 20.0)
	})

	t.Run("MinWidth might expand actual", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(10), Height(11), MinWidth(20), MinHeight(21))

		assert.Equal(t, box.Width(), 20.0)
		assert.Equal(t, box.Height(), 21.0)
	})

	t.Run("WidthInBounds", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(10), MaxWidth(20), Width(15))
		box.SetWidth(21)
		assert.Equal(t, box.Width(), 20.0)
		box.SetWidth(9)
		assert.Equal(t, box.Width(), 10.0)
		box.SetWidth(16)
		assert.Equal(t, box.Width(), 16.0)
	})

	t.Run("WidthInBounds from Child expansion plus Padding", func(t *testing.T) {
		box, err := Box(NewBuilder(), Padding(10), Width(30), Height(20), Children(func(b Builder) {
			Box(b, MinWidth(50), MinHeight(40))
			Box(b, MinWidth(30), MinHeight(30))
		}))

		if err != nil {
			t.Error(err)
			return
		}

		box.SetWidth(10)
		box.SetHeight(10)
		// This is a stack, so only the wider child expands parent.
		assert.Equal(t, box.Width(), 70.0)
		// assert.Equal(t, box.GetHeight(), 60.0)
	})

	t.Run("PrefWidth default value", func(t *testing.T) {
		one := NewComponent()
		assert.Equal(t, 0.0, one.PrefWidth())
	})

	t.Run("PrefWidth ComponentModel value", func(t *testing.T) {
		one, _ := TestComponent(NewBuilder(), PrefWidth(200))
		assert.Equal(t, 200.0, one.PrefWidth())
	})

	t.Run("GetOffsetFor", func(t *testing.T) {
		t.Run("Root at 0,0", func(t *testing.T) {
			root, _ := Box(NewBuilder())
			xOffset := root.XOffset()
			yOffset := root.YOffset()
			assert.Equal(t, xOffset, 0)
			assert.Equal(t, yOffset, 0)
		})

		t.Run("Root at offset", func(t *testing.T) {
			root, _ := Box(NewBuilder(), X(10), Y(15))
			xOffset := root.XOffset()
			yOffset := root.YOffset()
			assert.Equal(t, xOffset, 10)
			assert.Equal(t, yOffset, 15)
		})

		t.Run("Child receives offset for padding", func(t *testing.T) {
			var root, child Displayable
			root, _ = Box(NewBuilder(), Padding(10), Width(100), Height(100), Children(func(b Builder) {
				child, _ = Box(b, FlexWidth(1), FlexHeight(1))
			}))

			assert.Equal(t, root.XOffset(), 0)
			assert.Equal(t, child.XOffset(), 10)
		})

		t.Run("Child at double offset", func(t *testing.T) {
			var nestedChild Displayable
			Box(NewBuilder(), Padding(10), Children(func(b Builder) {
				Box(b, Padding(15), Children(func() {
					nestedChild, _ = Box(b, Padding(10))
				}))
			}))

			xOffset := nestedChild.XOffset()
			yOffset := nestedChild.YOffset()
			assert.Equal(t, xOffset, 25)
			assert.Equal(t, yOffset, 25)
		})
	})

	t.Run("Padding", func(t *testing.T) {
		t.Run("DefaultPadding", func(t *testing.T) {
			box, err := Box(NewBuilder())
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.Padding(), -1, "Default Padding")
			assert.Equal(t, box.PaddingBottom(), -1, "Default PaddingBottom")
			assert.Equal(t, box.PaddingTop(), -1, "Default PaddingTop")
			assert.Equal(t, box.PaddingLeft(), -1, "Default PaddingLeft")
			assert.Equal(t, box.PaddingRight(), -1, "Default PaddingRight")

			assert.Equal(t, box.MinWidth(), -1, "GetMinWidth")
			assert.Equal(t, box.MinHeight(), -1, "GetMinWidth")

			assert.Equal(t, box.Width(), 0, "Width")
		})

		t.Run("Override side padding", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.Padding(), 10, "Default Padding")
			assert.Equal(t, box.PaddingBottom(), 10, "Default PaddingBottom")
			assert.Equal(t, box.PaddingTop(), 10, "Default PaddingTop")
			assert.Equal(t, box.PaddingLeft(), 10, "Default PaddingLeft")
			assert.Equal(t, box.PaddingRight(), 10, "Default PaddingRight")
		})

		t.Run("Interacts with GetMinWidth()", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, box.MinWidth(), 20, "GetMinWidth")
			assert.Equal(t, box.MinHeight(), 20, "GetMinWidth")
		})

		t.Run("Applying Padding spreads to all four sides", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10))

			assert.Equal(t, root.HorizontalPadding(), 20.0)
			assert.Equal(t, root.VerticalPadding(), 20.0)

			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.PaddingLeft(), 10.0)
			assert.Equal(t, root.PaddingRight(), 10.0)
			assert.Equal(t, root.PaddingTop(), 10.0)
		})

		t.Run("PaddingTop overrides Padding", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10), PaddingTop(5))
			assert.Equal(t, root.PaddingTop(), 5.0)
			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.Padding(), 10.0)
		})

		t.Run("PaddingTop overrides Padding regardless of order", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), PaddingTop(5), Padding(10))
			assert.Equal(t, root.PaddingTop(), 5.0)
			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.Padding(), 10.0)
		})
	})
}
