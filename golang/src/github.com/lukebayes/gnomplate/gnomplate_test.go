package gnomplate

import "testing"

func TestGnomplate(t *testing.T) {
	instance := &Window{}
	if instance == nil {
		t.Errorf("Expected instance not nil: %v", instance)
	}
}
