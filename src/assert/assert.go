// Package assert includes helper functions that work with the native Go
// testing package.
package assert

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"testing"
)

func messagesToString(mainMessage string, optMessages ...string) (string, error) {
	switch len(optMessages) {
	case 0:
		return mainMessage, nil
	case 1:
		return fmt.Sprintf("%s\n%s", mainMessage, optMessages[0]), nil
	case 2:
		return fmt.Sprintf("%s\n%s\n%s", mainMessage, optMessages[0], optMessages[1]), nil
	default:
		return "", errors.New("Custom assertion provided with unexpected messages")

	}
}

// StrictEqual fails if the provided values are not == to one another.
func StrictEqual(t testing.TB, found interface{}, expected interface{}, message ...string) {
	if found != expected {
		mainMessage := fmt.Sprintf("Expected %v to STRICTLY equal %v", found, expected)
		msg, err := messagesToString(mainMessage, message...)
		if err != nil {
			t.Error(err)
			return
		}
		t.Error(msg)
	}
}

func float64EqualsInt(floatValue float64, intValue int) bool {
	intPart, div := math.Modf(floatValue)
	if div == 0.0 && int(intPart) == intValue {
		return true
	}
	return false
}

// Equal fails if the provided values are not equal in a "best effort" comparison.
// This method will (perhaps incorrectly to reasonably folks) claim 1.0 is
// equal to 1.
// This coercion helps with test brevity and flexibility. If you'd like
// something more precise, use StrictEqual instead.
func Equal(t testing.TB, found interface{}, expected interface{}, message ...string) {
	if found != expected {
		msg, msgErr := messagesToString("", message...)
		if msgErr != nil {
			t.Error(msgErr)
			return
		}
		kindA := reflect.ValueOf(found).Kind()
		kindB := reflect.ValueOf(expected).Kind()
		switch kindA {
		case reflect.Float64:
			if kindB == reflect.Int {
				if !float64EqualsInt(found.(float64), expected.(int)) {
					t.Errorf("Custom Equal expected %.2g to equal %v\n%s", found, expected, msg)
				}
				return
			}
		case reflect.Int:
			if kindB == reflect.Float64 {
				if !float64EqualsInt(expected.(float64), found.(int)) {
					t.Errorf("Custom Equal expected %.2g to equal %v\n%s", expected, found, msg)
				}
				return
			}
		}

		if found != expected {
			t.Errorf("Custom Equal expected %v to equal %v\n%s", found, expected, msg)
		}
	}
}

// Match fails if the the provided exprStr is not found in the provided str value as
// a regular expression.
func Match(t testing.TB, exprStr string, str string) {
	matched, _ := regexp.MatchString(exprStr, str)
	if !matched {
		t.Errorf("Expected: \"%v\", but received: \"%v\"", exprStr, str)
	}
}

func isTrue(t testing.TB, value bool, mainMessage string, message ...string) {
	if !value {
		msg, msgErr := messagesToString(mainMessage, message...)
		if msgErr != nil {
			t.Error(msgErr)
			return
		}
		t.Error(errors.New(msg).Error())
	}

}

// True fails if the provided value is not true
func True(t testing.TB, value bool, messages ...string) {
	isTrue(t, value, fmt.Sprintf("Expected %v to be true", value), messages...)
}

// False fails if the provided value is not false
func False(t testing.TB, value bool, messages ...string) {
	isTrue(t, !value, fmt.Sprintf("Expected %v to be false", value), messages...)
}

// NotNil fails if the provided value is nil
func NotNil(t testing.TB, value interface{}, messages ...string) {
	if value == nil {
		msg := fmt.Sprintf("Expected %v to not be nil", value)
		t.Errorf(messagesToString(msg, messages...))
	}
}

// Nil fails if the provided value is not nil
func Nil(t testing.TB, value interface{}, messages ...string) {
	if value != nil {
		typeOf := reflect.TypeOf(value).String()
		msg := fmt.Sprintf("Expected %v of type: %v to be nil", value, typeOf)
		t.Errorf(messagesToString(msg, messages...))
	}
}
