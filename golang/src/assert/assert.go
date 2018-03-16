package assert

import "fmt"

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
		panic(fmt.Errorf("Expected value to not be nil"))
	}
}

func Nil(value interface{}) {
	if value != nil {
		panic(fmt.Errorf("Expected value to be nil but was (%v)", value))
	}
}
