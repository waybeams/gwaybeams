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

func TStrictEqual(t testing.TB, found interface{}, expected interface{}, message ...string) {
	if found != expected {
		mainMessage := fmt.Sprintf("Custom Equal expected %v to strictly equal %v", found, expected)
		msg, err := messagesToString(mainMessage, message...)
		if err != nil {
			t.Error(err)
			return
		}
		t.Error(errors.New(msg))
	}

}

func float64EqualsInt(floatValue float64, intValue int) bool {
	intPart, div := math.Modf(floatValue)
	if div == 0.0 && int(intPart) == intValue {
		return true
	}
	return false
}

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

func True(t testing.TB, value bool, messages ...string) {
	isTrue(t, value, fmt.Sprintf("Expected %v to be true", value), messages...)
}

func False(t testing.TB, value bool, messages ...string) {
	isTrue(t, !value, fmt.Sprintf("Expected %v to be false", value), messages...)
}

func NotNil(t testing.TB, value interface{}, message ...string) {
}

func Nil(value interface{}) {
	if value != nil {
		panic(fmt.Errorf("Expected value to be nil but was (%v)", value))
	}
}
