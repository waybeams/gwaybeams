package controls

import (
	"assert"
	"events"
	"testing"
	"ui"
	"uiold/context"
	"uiold/opts"
)

func createLabel(text string) ui.Displayable {
	return Label(
		context.NewTestContext(),
		opts.FontFace("Roboto"),
		opts.FontSize(12),
		opts.Text(text),
	)
}

func TestLabel(t *testing.T) {
	t.Run("Label", func(t *testing.T) {
		label := createLabel("Hello World")
		assert.Equal(t, label.Text(), "Hello World")
		assert.Equal(t, label.Height(), 9, "MinHeight set")
		assert.Equal(t, label.Width(), 49, "MinWidth set")
	})

	t.Run("Metrics change when FontSize changes", func(t *testing.T) {
		t.Skip()
		label := createLabel("Hello")
		label.SetFontSize(36)
		label.Layout()
		label.Emit(events.New(events.Configured, label, nil))
		assert.Equal(t, label.Height(), 25, "MinHeight set")
		assert.Equal(t, label.Width(), 68, "MinWidth set")
	})
}
