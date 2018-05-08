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
		return fmt.Sprintf("%s %s", mainMessage, optMessages[0]), nil
	case 2:
		return fmt.Sprintf("%s %s\n%s", mainMessage, optMessages[0], optMessages[1]), nil
	default:
		return "", errors.New("Custom assertion provided with unexpected messages")

	}
}

// Panic fails if the provided handler does not trigger a panic that includes an error
// or message that matches the provided expression string.
func Panic(t testing.TB, expr string, handler func()) {
	defer func() {
		r := recover()
		if r != nil {
			var err error
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			Match(t, expr, err.Error())
		}
	}()
	handler()
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

func valueToKindAndString(value interface{}) (kind reflect.Kind, asString string) {
	kind = reflect.ValueOf(value).Kind()
	return kind, fmt.Sprintf("%v", value)
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
		switch kindA {
		case reflect.Bool:
			fallthrough
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			fallthrough
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			foundStr := fmt.Sprintf("%v", found)
			expectedStr := fmt.Sprintf("%v", expected)
			if foundStr != expectedStr {
				message := fmt.Sprintf("Custom Equal expected %v to equal %v. %s", found, expected, msg)
				t.Error(message)
			}
			return
		}

		if found != expected {
			t.Errorf("Custom Equal expected %v to equal %v. %s", found, expected, msg)
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
