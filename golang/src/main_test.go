package main

import "testing"

// NOTE: Added this test simply to prevent go test from complaining about "no
// test found"
func TestFake(t *testing.T) {
	if false {
		t.Fatal("Expected false to be true")
	}
}
