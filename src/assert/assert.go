package assert

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"testing"
)

func messageToString(message ...string) (string, error) {
	switch len(message) {
	case 0:
		return "", nil
	case 1:
		return message[0], nil
	default:
		return "", errors.New("Custom assertion provided with unexpected messages")

	}
}

func TStrictEqual(t testing.TB, found interface{}, expected interface{}, message ...string) {
	if found != expected {
		msg, err := messageToString(message...)
		if err != nil {
			t.Error(err)
			return
		}
		t.Errorf("Custom Equal expected %v to strictly equal %v\n%s", found, expected, msg)
	}

}

func float64EqualsInt(floatValue float64, intValue int) bool {
	intPart, div := math.Modf(floatValue)
	if div == 0.0 && int(intPart) == intValue {
		return true
	}
	return false
}

func TEqual(t testing.TB, found interface{}, expected interface{}, message ...string) {
	if found != expected {
		msg, msgErr := messageToString(message...)
		if msgErr != nil {
			t.Error(msgErr)
			return
		}
		typeA := reflect.TypeOf(found).String()
		typeB := reflect.TypeOf(expected).String()
		switch typeA {
		case "float64":
			if typeB == "int" {
				if !float64EqualsInt(found.(float64), expected.(int)) {
					t.Errorf("Custom Equal expected %.2g to equal %v\n%s", found, expected, msg)
				}
				return
			}
		case "int":
			if typeB == "float64" {
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

func TMatch(t testing.TB, exprStr string, str string) {
	matched, _ := regexp.MatchString(exprStr, str)
	if !matched {
		t.Errorf("Expected: \"%v\", but received: \"%v\"", exprStr, str)
	}
}

func Match(exprStr string, str string) {
	matched, _ := regexp.MatchString(exprStr, str)
	if !matched {
		panic(fmt.Errorf("Expected: \"%v\", but received: \"%v\"", exprStr, str))
	}
}

func ErrorMatch(exprStr string, err error) {
	if err == nil {
		panic(errors.New("Expected error response"))
	}
	Match(exprStr, err.Error())
}

func True(value bool) {
	if !value {
		panic(fmt.Errorf("Expected %t to be true", value))
	}
}

func False(value bool) {
	if value {
		panic(fmt.Errorf("Expected %t to be false", value))
	}
}

func Equal(value interface{}, expected interface{}) {
	if value != expected {
		panic(fmt.Errorf("Expected (%v) to equal (%v)", value, expected))
	}
}

func NotEqual(value interface{}, expected interface{}) {
	if value == expected {
		panic(fmt.Errorf("Expected (%v) to not equal (%v)", value, expected))
	}
}

func NotNil(value interface{}) {
	if value == nil {
		panic(errors.New("Expected value to not be nil"))
	}
}

func Nil(value interface{}) {
	if value != nil {
		panic(fmt.Errorf("Expected value to be nil but was (%v)", value))
	}
}
