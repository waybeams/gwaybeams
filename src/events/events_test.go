package events

import "testing"

func TestEventNames(t *testing.T) {
	t.Run("No duplicate values", func(t *testing.T) {
		for index, eventName := range AllEvents {
			for k, otherName := range AllEvents {
				if index != k && eventName == otherName {
					t.Errorf("Duplicate event value found %v", eventName)
				}
			}
		}
	})
}
