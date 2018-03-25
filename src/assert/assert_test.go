package assert

import (
	"fmt"
	"reflect"
	"testing"
)

type CustomT struct {
	testing.T
	failureMsg string
}

func (c *CustomT) Errorf(format string, args ...interface{}) {
	c.failureMsg = fmt.Sprintf(format, args...)
}

func (c *CustomT) Error(msgOrErr ...interface{}) {
	msg := msgOrErr[0]
	msgType := reflect.TypeOf(msg).String()
	fmt.Println("msg:", msg)

	switch msgType {
	case "string":
		c.failureMsg = msg.(string)
	case "error":
		c.failureMsg = msg.(error).Error()
	default:
		panicMsg := fmt.Sprintf("Unexpected call to CustomT.Error with type: %s", msgType)
		panic(panicMsg)
	}
}

func NewCustomT() *CustomT {
	return &CustomT{}
}

func TestAssertions(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ct := NewCustomT()
			Match(ct, "foo", "sdffoosdf")
			if ct.failureMsg != "" {
				t.Errorf("Unexpected failure %s", ct.failureMsg)
			}
		})

		t.Run("Failure message", func(t *testing.T) {
			ct := NewCustomT()
			Match(ct, "foo", "sdf")
			if ct.failureMsg != "Expected: \"foo\", but received: \"sdf\"" {
				t.Error(ct.failureMsg)

			}
		})
	})

	t.Run("NotNil", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ct := NewCustomT()
			NotNil(ct, true)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}

		})

		t.Run("Failure message", func(t *testing.T) {
			ct := NewCustomT()
			NotNil(ct, nil)
			if ct.failureMsg == "" {
				t.Error("Expected failure")
			}
			Match(t, "not be nil", ct.failureMsg)
		})
	})

	t.Run("Nil", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ct := NewCustomT()
			Nil(ct, nil)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

	})

	t.Run("True", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ct := NewCustomT()
			True(ct, true)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})
	})

	t.Run("False", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ct := NewCustomT()
			False(ct, false)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

		t.Run("Failure message", func(t *testing.T) {
			ct := NewCustomT()
			False(ct, true)
			if ct.failureMsg == "" {
				t.Error("Expected a failure message")
			}
			Match(t, "Expected true to be false", ct.failureMsg)
		})
	})

	t.Run("Equality helper", func(t *testing.T) {
		t.Run("0.0 == 0.0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0.0, 0.0)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

		t.Run("0.0 == 0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0.0, 0)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

		t.Run("0 == 0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0, 0)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

		t.Run("0 == 0.0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0, 0.0)
			if ct.failureMsg != "" {
				t.Error(ct.failureMsg)
			}
		})

		t.Run("failure with custom message", func(t *testing.T) {
			ct := NewCustomT()

			Equal(ct, 1, 2, "Fake custom message")
			Match(t, "Fake custom message", ct.failureMsg)
		})
	})
}

// Stop calling panic in assertions!
// TODO(lbayes): Add tests for failing assertions when we stop calling PANIC
// whenever a test fails.
