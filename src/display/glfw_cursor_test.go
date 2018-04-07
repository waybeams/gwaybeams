package display

import (
	"testing"
)

func TestCursor(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		fakeWindow := NewFakeWindow()
		cursor := NewCursor(fakeWindow)
		if cursor == nil {
			t.Error("Expected instance of Cursor")
		}
	})
}
