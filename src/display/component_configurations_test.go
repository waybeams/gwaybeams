package display

import (
	"assert"
	"log"
	"testing"
)

func TestVBox(t *testing.T) {

	t.Run("Simple Children", func(t *testing.T) {
		var err error

		root, _ := VBox(NewBuilder(), Height(100), Children(func(b Builder) {
			_, err = Box(b, FlexHeight(1))
			Box(b, FlexHeight(1))
		}))

		root.Layout()

		one := root.GetChildAt(0)
		two := root.GetChildAt(1)
		if err != nil {
			t.Error(err)
		}
		if one.GetHeight() != 50 {
			t.Errorf("one expected 50, but was %v", one.GetHeight())
		}
		if two.GetHeight() != 50 {
			t.Errorf("two expected 50, but was %v", one.GetHeight())
		}
	})
}

func TestLabel(t *testing.T) {
	t.Run("Label", func(t *testing.T) {
		label, _ := Label(NewBuilder(), Title("Hello World"))
		assert.Equal(t, label.GetTitle(), "Hello World")
	})
}

func TestButton(t *testing.T) {
	t.Run("Clickable", func(t *testing.T) {
		var clickHandler = func(btn Displayable) {
			log.Println("CLIIIIIIIIIIIIIIIIIIIICK")
		}
		button, _ := Box(NewBuilder(), Text("Submit"), OnClick(clickHandler))
		button.Click()
	})
}
