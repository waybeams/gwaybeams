package assert

import (
	"testing"
)

func TestSuccessAssertions(t *testing.T) {
	Nil(nil)
	NotNil(true)
	True(true)
	False(false)
	Equal(0, 0)
	NotEqual(0, 1)
}

// Stop calling panic in assertions!
// TODO(lbayes): Add tests for failing assertions when we stop calling PANIC
// whenever a test fails.
