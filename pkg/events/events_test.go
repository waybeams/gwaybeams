package events_test

import (
	"github.com/waybeams/waybeams/pkg/events"
	"testing"
)

func TestEventNames(t *testing.T) {
	t.Run("No duplicate values", func(t *testing.T) {
		for index, eventName := range events.AllEvents {
			for k, otherName := range events.AllEvents {
				if index != k && eventName == otherName {
					t.Errorf("Duplicate event value found %v", eventName)
				}
			}
		}
	})
}
