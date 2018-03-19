package assert

import (
	"errors"
	"fmt"
	"regexp"
)

func ErrorMatches(exprStr string, err error) {
	if err == nil {
		panic(errors.New("Expected error response"))
	}
	errStr := err.Error()
	matched, _ := regexp.MatchString(exprStr, errStr)
	if !matched {
		panic(fmt.Errorf("Expected: \"%v\", but received: \"%v\"", exprStr, errStr))
	}
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
