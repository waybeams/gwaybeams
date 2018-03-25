package assert

import (
	"fmt"
	"reflect"
	"testing"
)

type CustomT struct {
	testing.T
	failureWith string
}

func (c *CustomT) Errorf(format string, args ...interface{}) {
	c.failureWith = fmt.Sprintf(format, args...)
}

func (c *CustomT) Error(msgOrErr ...interface{}) {
	fmt.Println("ERROR CALLED ITH:", msgOrErr)
	msg := msgOrErr[0]
	msgType := reflect.TypeOf(msg).String()
	fmt.Println("msg:", msg)

	switch msgType {
	case "string":
		c.failureWith = msg.(string)
	case "error":
		c.failureWith = msg.(error).Error()
	default:
		panic("Unexpected call to CustomT.Error")
	}
}

func NewCustomT() *CustomT {
	return &CustomT{}
}

type assertCase struct {
	failureExpr string
	found       interface{}
	expected    interface{}
}

func TestSuccessAssertions(t *testing.T) {
	Nil(nil)
	NotNil(true)
	True(true)
	False(false)
	NotEqual(0, 1)

	t.Run("Equality helper", func(t *testing.T) {
		t.Run("0.0 == 0.0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0.0, 0.0)
			if ct.failureWith != "" {
				t.Error(ct.failureWith)
			}
		})

		t.Run("0.0 == 0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0.0, 0)
			if ct.failureWith != "" {
				t.Error(ct.failureWith)
			}
		})

		t.Run("0 == 0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0, 0)
			if ct.failureWith != "" {
				t.Error(ct.failureWith)
			}
		})

		t.Run("0 == 0.0", func(t *testing.T) {
			ct := NewCustomT()
			Equal(ct, 0, 0.0)
			if ct.failureWith != "" {
				t.Error(ct.failureWith)
			}
		})

		t.Run("failure with custom message", func(t *testing.T) {
			ct := NewCustomT()

			Equal(ct, 1, 2, "Fake custom message")
			TMatch(t, "Fake custom message", ct.failureWith)
		})
	})
}

// Stop calling panic in assertions!
// TODO(lbayes): Add tests for failing assertions when we stop calling PANIC
// whenever a test fails.
