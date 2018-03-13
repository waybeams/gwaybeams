package gnomplate

import "testing"

func TestGnomplate(t *testing.T) {
	instance := New()
	if instance == nil {
		t.Fatal("Expected instance of Gnomplate")
	}
}
