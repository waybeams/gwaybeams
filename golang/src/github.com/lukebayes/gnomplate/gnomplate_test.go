package gnomplate

import "testing"

func TestGnomplate(t *testing.T) {
	instance := New()
	if instance == nil {
		t.Errorf("Expected instance not nil: %s", instance)
	}
}
